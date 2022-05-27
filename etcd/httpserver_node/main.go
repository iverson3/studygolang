package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"sync"
	"time"
)

// 签到

const TIME_LAYOUT = "2006-01-02 15:04:05"

// SignMetaData 用户签到相关元数据
type SignMetaData struct {
	UserId string
	ContinueSignStartDate string  // 连续签到开启日期
	ContinueSignSchedule string  // 连续签到任务完成进度 json [1,0,0]
	StageSignStartDate string  // 阶段签到开始时间<1.每次成功发放连续签到的额外积分需要清空该值，2.满足周期重置时间也需要清空>
	CycleSignStartDate string  // 周期累计签到开始时间<1.每次成功发放连续签到的额外积分需要清空该值，2.满足周期重置时间也需要清空>
}
// DailySignCfg 每日打卡奖励配置
type DailySignCfg struct {
	IsEnabled int
	Type int
	FixedCfg
	StepCfg
}
type FixedCfg struct {
	Score int
}
type StepCfg struct {
	IsReset bool
	ResetCycle int
	Rules []StepRule
}
type StepRule struct {
	Days int
	Score int
}

// ContinueSignCfg 连续打卡奖励配置
type ContinueSignCfg struct {
	IsEnabled int
	ContinueCfg
}
type ContinueCfg struct {
	IsReset bool
	ResetCycle int
	Rules []StepRule
}

// CycleSignTotalCfg 周期累计打卡奖励配置
type CycleSignTotalCfg struct {
	IsEnabled int
	Type int
	CycleCfg
}
type CycleCfg struct {
	ResetCycleDays int
	Days int
	Score int
}

// Sign 打卡签到实现
func Sign(curDate string) {
	// 从缓存或数据库中获取用户签到相关元数据
	signMetaData := SignMetaData{
		UserId:                "101",
		ContinueSignStartDate: "2022-03-26",
		ContinueSignSchedule:  "[1,0,0]",
		StageSignStartDate:    "2022-03-26",
		CycleSignStartDate:    "2022-03-24",
	}
	fmt.Println(signMetaData)

	// 从缓存或数据库中获取配置信息
	dailySignCfg := DailySignCfg{
		IsEnabled: 1,
		Type:      2,
		FixedCfg:  FixedCfg{Score: 3},
		StepCfg:   StepCfg{
			IsReset:    true,
			ResetCycle: 14,
			Rules:      []StepRule{
				{
					Days:  0,
					Score: 2,
				},
				{
					Days:  5,
					Score: 5,
				},
				{
					Days:  10,
					Score: 12,
				},
			},
		},
	}
	continueSignCfg := ContinueSignCfg{
		IsEnabled:   1,
		ContinueCfg: ContinueCfg{
			IsReset:    true,
			ResetCycle: 14,
			Rules:      []StepRule{
				{
					Days:  3,
					Score: 5,
				},
				{
					Days:  9,
					Score: 12,
				},
				{
					Days:  14,
					Score: 20,
				},
			},
		},
	}
	cycleSignTotalCfg := CycleSignTotalCfg{
		IsEnabled: 1,
		Type:      1,
		CycleCfg:  CycleCfg{
			ResetCycleDays: 14,
			Days:           9,
			Score:          25,
		},
	}
	fmt.Println(dailySignCfg)
	fmt.Println(continueSignCfg)
	fmt.Println(cycleSignTotalCfg)

	wg := sync.WaitGroup{}
	var dailySignIntegral int    // 每日签到积分
	var continueSignIntegral int // 连续签到积分
	var totalSignIntegral int    // 累计签到积分
	// 协程同时处理三种签到
	wg.Add(3)
	start := time.Now()
	go DailySign(&wg, curDate, signMetaData, dailySignCfg, &dailySignIntegral)
	go ContinueSign(&wg, curDate, signMetaData, continueSignCfg, &continueSignIntegral)
	go TotalSign(&wg, curDate, signMetaData, cycleSignTotalCfg, &totalSignIntegral)

	wg.Wait()
	duration := time.Since(start).Nanoseconds()
	fmt.Println(dailySignIntegral, continueSignIntegral, totalSignIntegral)
	fmt.Println("总积分：", dailySignIntegral + continueSignIntegral + totalSignIntegral)
	fmt.Println("耗时：", duration ,"纳秒")
}

// DailySign 每日签到
func DailySign(wg *sync.WaitGroup, curDate string, signMetaData SignMetaData, cfg DailySignCfg, integral *int) {
	defer func() {
		// todo: 将bitmap指定位置设置为 1 表示今天已打卡
		wg.Done()
	}()
	// 是否开启了配置
	if cfg.IsEnabled == 0 {
		return
	}

	if cfg.Type == 1 {
		// 每日固定奖励处理
		*integral = cfg.FixedCfg.Score
		return
	}

	// 阶梯奖励配置处理
	// 判断“阶段签到开始日期”是否为空
	if signMetaData.StageSignStartDate == "" {
		// todo: 设置当前日期为“阶段签到开始日期”
		*integral = cfg.StepCfg.Rules[0].Score
		return
	}

	curDateParse, _ := time.Parse(TIME_LAYOUT, fmt.Sprintf("%s 00:00:00", curDate))
	stageSignStartDate, _ := time.Parse(TIME_LAYOUT, fmt.Sprintf("%s 00:00:00", signMetaData.StageSignStartDate))

	// 判断前一天是否已打卡
	signed := yesterdayIsSigned(curDateParse)
	if signed {
		goto End
	}

	// 代码能执行到这里，说明前一天没有打卡
	// todo: 设置当前日期为“阶段签到开始日期”
	*integral = cfg.StepCfg.Rules[0].Score
	return

End:
	// 计算阶梯连签天数
	stageContinueSignDays := int((curDateParse.Sub(stageSignStartDate).Hours() / 24) + 1)
	fmt.Println("stageContinueSignDays", stageContinueSignDays)

	// 判断是否开启周期重置
	if cfg.StepCfg.IsReset {
		if stageContinueSignDays >= cfg.StepCfg.ResetCycle {
			if stageContinueSignDays == cfg.StepCfg.ResetCycle {
				// todo: 置空“阶段签到开始日期”
			} else {
				// todo: 设置当前日期为“阶段签到开始日期”
				*integral = cfg.StepCfg.Rules[0].Score
				return
			}
		}
	}

	for i := len(cfg.StepCfg.Rules) - 1; i >= 0; i-- {
		if stageContinueSignDays >= cfg.StepCfg.Rules[i].Days {
			*integral = cfg.StepCfg.Rules[i].Score
			return
		}
	}
}

// ContinueSign 连续签到
func ContinueSign(wg *sync.WaitGroup, curDate string, signMetaData SignMetaData, cfg ContinueSignCfg, integral *int) {
	defer func() {
		// todo: 将bitmap指定位置设置为 1 表示今天已打卡
		wg.Done()
	}()

	// 是否开启了配置
	if cfg.IsEnabled == 0 {
		return
	}

	var continueSignDays int   // 连签天数
	// 判断“连续签到开始日期”是否为空
	if signMetaData.ContinueSignStartDate == "" {
		// todo: 设置当前日期为“阶段签到开始日期”
		continueSignDays = 1
	} else {
		curDateParse, _ := time.Parse(TIME_LAYOUT, fmt.Sprintf("%s 00:00:00", curDate))
		continueSignStartDate, _ := time.Parse(TIME_LAYOUT, fmt.Sprintf("%s 00:00:00", signMetaData.ContinueSignStartDate))

		// 判断前一天是否已打卡
		signed := yesterdayIsSigned(curDateParse)
		if signed {
			// 计算阶段连签天数
			continueSignDays = int((curDateParse.Sub(continueSignStartDate).Hours() / 24) + 1)
			fmt.Println("continueSignDays", continueSignDays)
		} else {
			// 前一天没有打卡
			// todo: 设置当前日期为“连续签到开始日期”
			continueSignDays = 1
		}
	}

	if !cfg.ContinueCfg.IsReset {
		goto End
	}

	// 开启了周期重置，则判断连签天数和重置周期的关系
	if continueSignDays < cfg.ContinueCfg.ResetCycle {
		goto End
	}
	if continueSignDays == cfg.ContinueCfg.ResetCycle {
		// todo: 置空“连续签到开始日期”
		goto End
	}
	if continueSignDays > cfg.ContinueCfg.ResetCycle {
		// todo: 设置当前日期为“连续签到开始日期”
		continueSignDays = 1
		goto End
	}

End:
	// 获取该用户连签任务进度信息，开始计算
	continueSignSchedule := make([]int64, 0)
	if err := json.Unmarshal([]byte(signMetaData.ContinueSignSchedule), &continueSignSchedule); err != nil {
		panic(err)
	}

	for i := len(cfg.ContinueCfg.Rules) - 1; i >= 0; i-- {
		if continueSignDays >= cfg.ContinueCfg.Rules[i].Days {
			if len(continueSignSchedule) > i && continueSignSchedule[i] == 0 {
				*integral = cfg.ContinueCfg.Rules[i].Score
				// 更新该用户的连签任务进度信息
				continueSignSchedule[i] = 1
			}
			return
		}
	}
}

// TotalSign 累计签到
func TotalSign(wg *sync.WaitGroup, curDate string, signMetaData SignMetaData, cfg CycleSignTotalCfg, integral *int)  {
	defer func() {
		// todo: 将bitmap指定位置设置为 1 表示今天已打卡
		wg.Done()
	}()

	// 是否开启了配置
	if cfg.IsEnabled == 0 {
		return
	}

	var totalSignDays int   // 累计签到天数
	var n int  // 用户在累计签到配置的一个有效周期内的所经天数
	// 判断“累计签到开始日期”是否为空
	if signMetaData.CycleSignStartDate == "" {
		// todo: 设置当前日期为“累计签到开始日期”
		totalSignDays = 1
		goto End
	} else {
		curDateParse, _ := time.Parse(TIME_LAYOUT, fmt.Sprintf("%s 00:00:00", curDate))
		totalSignStartDate, _ := time.Parse(TIME_LAYOUT, fmt.Sprintf("%s 00:00:00", signMetaData.CycleSignStartDate))

		// 计算从"累计签到开启日期"到今天所经过的天数
		totalSignDuration := int((curDateParse.Sub(totalSignStartDate).Hours() / 24) + 1)
		fmt.Println("totalSignDuration", totalSignDuration)

		x := totalSignDuration % cfg.CycleCfg.ResetCycleDays
		if x == 0 {
			n = cfg.CycleCfg.ResetCycleDays
		} else {
			n = x
		}

		days := curDateParse.Day()
		curMonthDays := CalDaysOnMonth(curDateParse.Year(), int(curDateParse.Month()))
		// 当天的bitmap还没设置为1，初始totalSignDays = 1 即提前算上了这一天的签到天数
		totalSignDays = 1
		if n > days {
			lastMonthDuration := n - days
			// 需要结合上月的bitmap一起计算累计签到天数
			// 利用bitmap计算从当前日期到往前n天之间的累计签到天数
			// 当月部分统计签到天数
			// 当天的bitmap还没设置为1 通过days--直接跳过对这一位是否为1的判断
			days--
			for days != 0 {
				// 将指定位数的二进制位 置为1，然后与自己相比; 相等则表示该位为1
				if (bitMap4 | (1 << (curMonthDays - days))) == bitMap4 {
					totalSignDays++
				}
				days--
			}
			// 上月部分统计签到天数
			for i := 0; i < lastMonthDuration; i++ {
				// 将指定位数的二进制位 置为1，然后与自己相比; 相等则表示该位为1
				if (bitMap3 | (1 << i)) == bitMap3 {
					totalSignDays++
				}
			}
		} else {
			// 只需使用当月的bitmap来计算累计签到天数
			// 利用bitmap计算从当前日期到往前n天之间的累计签到天数
			n--
			for n != 0 {
				// 将指定位数的二进制位 置为1，然后与自己相比; 相等则表示该位为1
				if (bitMap4 | (1 << (curMonthDays - n))) == bitMap4 {
					totalSignDays++
				}
				n--
			}
		}
		fmt.Println("totalSignDays", totalSignDays)
		goto End
	}

End:
	if totalSignDays >= cfg.CycleCfg.Days {
		*integral = cfg.CycleCfg.Score
	}
	if n == cfg.CycleCfg.ResetCycleDays {
		// todo: 设置"累计签到开启日期"为空
	}
}

// 判断前一天是否已打卡
func yesterdayIsSigned(curDate time.Time) bool {
	beforeOneDay, _ := time.ParseDuration("-24h")  // 往前倒一天
	yesterday := curDate.Add(beforeOneDay)

	// 前一天是上个月，即跨月了
	if yesterday.Month() == curDate.Month() - 1 {
		// 获取上一个月的bitmap - bitMap3
		// 将指定位数的二进制位 置为1，然后与自己相比 (将右边最后一位置为 1)
		b3 := bitMap3 | (1 << 0)
		// 相等则说明前一天是1 已打卡
		if b3 == bitMap3 {
			return true
		}
	} else {
		// 否则 没有跨月
		// 计算当月的总天数
		curMonthDays := CalDaysOnMonth(curDate.Year(), int(curDate.Month()))
		// 获取今天是当月的第几天
		n := curDate.Day()
		// 获取当月的bitmap - bitMap4
		// 将指定位数的二进制位 置为1，然后与自己相比
		b4 := bitMap4 | (1 << (curMonthDays - n + 1))
		// 相等则说明前一天是1 已打卡
		if b4 == bitMap4 {
			return true
		}
	}
	return false
}

// CalDaysOnMonth 计算指定年月份有多少天
func CalDaysOnMonth(year int, month int) int {
	// 有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] == true {
		return 31
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 闰年
		return 29
	}
	// 平年
	return 28
}

// CalSignRewardOnMonth 计算当月的预签到奖励
func CalSignRewardOnMonth(curDate string) {
	curDateParse, _ := time.Parse(TIME_LAYOUT, fmt.Sprintf("%s 00:00:00", curDate))
	beforeOneDay, _ := time.ParseDuration("-24h")  // 往前倒一天
	yesterday := curDateParse.Add(beforeOneDay)

	isSigned := yesterdayIsSigned(curDateParse)
	// 昨天和今天是同一个月份 且昨天打卡了
	if yesterday.Month() == curDateParse.Month() && isSigned {
		// todo: 直接返回之前计算出来并放在缓存中的月签到奖励预计算信息
		return
	}

	// 获取当月的总天数
	totalDays := CalDaysOnMonth(curDateParse.Year(), int(curDateParse.Month()))
	// 获取今天是本月第几天
	today := curDateParse.Day()

	// 从缓存或数据库中获取用户签到相关元数据
	//signMetaData := SignMetaData{
	//	UserId:                "101",
	//	ContinueSignStartDate: "2022-03-26",
	//	ContinueSignSchedule:  "[1,1,0]",
	//	StageSignStartDate:    "2022-03-26",
	//	CycleSignStartDate:    "2022-03-26",
	//}
	signMetaData := SignMetaData{
		UserId:                "101",
		ContinueSignStartDate: "2022-04-03",
		ContinueSignSchedule:  "[0,0,0]",
		StageSignStartDate:    "2022-04-03",
		CycleSignStartDate:    "2022-03-26",
	}
	//signMetaData := SignMetaData{
	//	UserId:                "101",
	//	ContinueSignStartDate: "",
	//	ContinueSignSchedule:  "",
	//	StageSignStartDate:    "",
	//	CycleSignStartDate:    "",
	//}
	fmt.Println(signMetaData)

	// 从缓存或数据库中获取配置信息
	dailySignCfg := DailySignCfg{
		IsEnabled: 1,
		Type:      2,
		FixedCfg:  FixedCfg{Score: 3},
		StepCfg:   StepCfg{
			IsReset:    true,
			ResetCycle: 14,
			Rules:      []StepRule{
				{
					Days:  0,
					Score: 2,
				},
				{
					Days:  5,
					Score: 5,
				},
				{
					Days:  10,
					Score: 12,
				},
			},
		},
	}
	continueSignCfg := ContinueSignCfg{
		IsEnabled:   1,
		ContinueCfg: ContinueCfg{
			IsReset:    true,
			ResetCycle: 14,
			Rules:      []StepRule{
				{
					Days:  3,
					Score: 5,
				},
				{
					Days:  9,
					Score: 12,
				},
				{
					Days:  14,
					Score: 20,
				},
			},
		},
	}
	cycleSignTotalCfg := CycleSignTotalCfg{
		IsEnabled: 1,
		Type:      1,
		CycleCfg:  CycleCfg{
			ResetCycleDays: 14,
			Days:           9,
			Score:          25,
		},
	}
	fmt.Println("dailySignCfg", dailySignCfg)
	fmt.Println("continueSignCfg", continueSignCfg)
	fmt.Println("cycleSignTotalCfg", cycleSignTotalCfg)

	// 初始化几个循环中使用的关键变量
	// 阶梯连续签到天数
	var stageContinueSignDays int
	if dailySignCfg.IsEnabled == 1 && dailySignCfg.Type == 2 {
		if isSigned && signMetaData.StageSignStartDate != "" {
			stageSignStartDate, _ := time.Parse(TIME_LAYOUT, fmt.Sprintf("%s 00:00:00", signMetaData.StageSignStartDate))
			stageContinueSignDays = int(curDateParse.Sub(stageSignStartDate).Hours() / 24)
		} else {
			stageContinueSignDays = 0
		}
	}
	// 连续签到天数
	var continueSignDays int
	if continueSignCfg.IsEnabled == 1 {
		if isSigned && signMetaData.ContinueSignStartDate != "" {
			continueSignStartDate, _ := time.Parse(TIME_LAYOUT, fmt.Sprintf("%s 00:00:00", signMetaData.ContinueSignStartDate))
			continueSignDays = int(curDateParse.Sub(continueSignStartDate).Hours() / 24)
		} else {
			continueSignDays = 0
		}
	}
	// 当前日期最近一个累计签到周期内的有效天数 (不是累计签到天数) (比如今天是4.5号，累计签到周期是14天,累计签到周期起始日期是3.26，那么totalSignDays=10)
	var totalSignDays int
	// 临时变量记录累计签到的周期起始日期距离月初的天数 (比如起始日期是4.5，那么distance=4)
	var distance int
	// 累计签到的最近一个周期中，位于上个月部分的天数 (比如今天是4.5号，累计签到周期是14天,累计签到周期起始日期是3.26，那么daysLastMonth=6)
	var daysLastMonth int
	if cycleSignTotalCfg.IsEnabled == 1 {
		if signMetaData.CycleSignStartDate == "" {
			totalSignDays = 0
			distance = today - 1
		} else {
			totalSignStartDate, _ := time.Parse(TIME_LAYOUT, fmt.Sprintf("%s 00:00:00", signMetaData.CycleSignStartDate))
			m := int(curDateParse.Sub(totalSignStartDate).Hours() / 24)
			if x := m % cycleSignTotalCfg.CycleCfg.ResetCycleDays; x == 0 {
				totalSignDays = 0
				distance = today - 1
			} else {
				totalSignDays = x
				// 累计签到周期起始日期是否位于上个月
				if x > today {
					distance = 0
					daysLastMonth = x - today + 1
				} else {
					distance = today - x - 1
				}
			}
		}
	}
	// 连续签到任务完成进度信息
	continueSignSchedule := make([]int, 0)
	// todo: 判断缓存和数据库中是否存在连签任务进度信息
	if signMetaData.ContinueSignSchedule == "" {
		// 不存在则根据连签奖励配置来初始化任务进度信息
		for _ = range continueSignCfg.ContinueCfg.Rules {
			continueSignSchedule = append(continueSignSchedule, 0)
		}
	} else {
		if err := json.Unmarshal([]byte(signMetaData.ContinueSignSchedule), &continueSignSchedule); err != nil {
			panic(err)
		}
	}

	// 预计算积分奖励数据
	preCalRewardList := make([]int, totalDays + 1)

	var dailySignIntegral, stageContinueSignIntegral, continueSignIntegral, totalSignIntegral int
	var i int
	// 从当天遍历到月底，依次计算每天需要预发放的积分
	for day := today; day <= totalDays; day++ {
		// 计算出当前天的日期
		//afterIDay, _ := time.ParseDuration(fmt.Sprintf("+%dh", 24 * i))  // 往后加i天
		//curDay := curDateParse.Add(afterIDay)

		// 清空积分数
		dailySignIntegral = 0
		stageContinueSignIntegral = 0
		continueSignIntegral = 0
		totalSignIntegral = 0
		// 每循环一遍就增加一天的天数
		i++
		stageContinueSignDays++
		continueSignDays++
		totalSignDays++

		// 每日签到奖励相关计算 (包括固定奖励和阶梯奖励)
		// 先判断是否开启配置
		if dailySignCfg.IsEnabled == 1 {
			// 阶梯奖励
			if dailySignCfg.Type == 2 {
				// 是否开启周期重置
				if dailySignCfg.StepCfg.IsReset {
					if stageContinueSignDays < dailySignCfg.StepCfg.ResetCycle {
						for j := len(dailySignCfg.StepCfg.Rules) - 1; j >= 0; j-- {
							if stageContinueSignDays >= dailySignCfg.StepCfg.Rules[j].Days {
								stageContinueSignIntegral = dailySignCfg.StepCfg.Rules[j].Score
								break
							}
						}
					} else if stageContinueSignDays == dailySignCfg.StepCfg.ResetCycle {
						for j := len(dailySignCfg.StepCfg.Rules) - 1; j >= 0; j-- {
							if stageContinueSignDays >= dailySignCfg.StepCfg.Rules[j].Days {
								stageContinueSignIntegral = dailySignCfg.StepCfg.Rules[j].Score
								break
							}
						}
						stageContinueSignDays = 0
					} else {

					}
				} else {
					for j := len(dailySignCfg.StepCfg.Rules) - 1; j >= 0; j-- {
						if stageContinueSignDays >= dailySignCfg.StepCfg.Rules[j].Days {
							stageContinueSignIntegral = dailySignCfg.StepCfg.Rules[j].Score
							break
						}
					}
				}
			} else {
				// 固定奖励
				dailySignIntegral = dailySignCfg.FixedCfg.Score
			}
		}

		// 连续签到奖励相关计算
		// 先判断是否开启配置
		if continueSignCfg.IsEnabled == 1 {
			var shouldReset bool
			if continueSignCfg.IsReset {
				if continueSignDays < continueSignCfg.ResetCycle {
					
				} else if continueSignDays == continueSignCfg.ResetCycle {
					shouldReset = true
				} else {
					
				}
			} else {
				
			}

			// 计算
			for j := len(continueSignCfg.ContinueCfg.Rules) - 1; j >= 0; j-- {
				if continueSignDays == continueSignCfg.ContinueCfg.Rules[j].Days {
					// 判断是否已经领取过
					if j < len(continueSignSchedule) && continueSignSchedule[j] == 0 {
						continueSignIntegral = continueSignCfg.ContinueCfg.Rules[j].Score
						// 标记该连续签到天数的额外奖励已经领取过了
						continueSignSchedule[j] = 1
					}
					break
				}
			}
			
			if shouldReset {
				continueSignDays = 0
				// 随着连续签到周期重置，同时重置连续签到的任务进度信息
				for k := range continueSignSchedule {
					continueSignSchedule[k] = 0
				}
			}
		}

		// 累计签到奖励相关计算
		// 先判断是否开启配置
		if cycleSignTotalCfg.IsEnabled == 1 {
			// 统计累计签到天数
			var n int
			days := day
			totalSignDuration := totalSignDays
			// 当天的bitmap还没设置为1，初始n = 1 即提前算上了这一天的签到天数
			n = 1
			// 需要结合两个月的bitmap来计算累计签到天数
			if totalSignDuration > days {
				// 上个月需要计算的bitmap位数
				lastMonthDuration := totalSignDuration - days
				// 需要结合上月的bitmap一起计算累计签到天数
				// 利用bitmap计算从当前日期到往前n天之间的累计签到天数
				// 当月部分统计签到天数
				// 当天的bitmap还没设置为1 通过days--直接跳过对这一位是否为1的判断
				days--
				for days != 0 {
					// 将指定位数的二进制位 置为1，然后与自己相比; 相等则表示该位为1
					if (bitMap4 | (1 << (totalDays - days))) == bitMap4 {
						n++
					}
					days--
				}
				// 上月部分统计签到天数
				for i := 0; i < lastMonthDuration; i++ {
					// 将指定位数的二进制位 置为1，然后与自己相比; 相等则表示该位为1
					if (bitMap3 | (1 << i)) == bitMap3 {
						n++
					}
				}
			} else {
				// 只需使用当月的bitmap来计算累计签到天数
				// 利用bitmap计算从当前日期到往前n天之间的累计签到天数
				totalSignDuration--
				for totalSignDuration != 0 {
					// 将指定位数的二进制位 置为1，然后与自己相比; 相等则表示该位为1
					if (bitMap4 | (1 << (totalDays - distance - totalSignDuration))) == bitMap4 {
						n++
					}
					totalSignDuration--
				}
			}

			if n == cycleSignTotalCfg.CycleCfg.Days {
				totalSignIntegral = cycleSignTotalCfg.CycleCfg.Score
			}

			if totalSignDays == cycleSignTotalCfg.CycleCfg.ResetCycleDays {
				totalSignDays = 0
				distance = distance + cycleSignTotalCfg.CycleCfg.ResetCycleDays - daysLastMonth
				if daysLastMonth != 0 {
					daysLastMonth = 0
				}
			}
		}

		// 将bitmap指定位数的二进制位 置为1，即表示今天已签到
		bitMap4 = bitMap4 | (1 << (totalDays - day))

		// 计算当日的总积分
		totalIntegral := dailySignIntegral + stageContinueSignIntegral + continueSignIntegral + totalSignIntegral
		preCalRewardList[day] = totalIntegral
	}

	fmt.Println(preCalRewardList)
}

// 记录用户签到情况的bitmap (以月为单位)  1-已签到 0-未签到

// 3月  1100111111100000111111100111111     31天
//      1 1 0 0 1 1 1 1 1 1  1  0  0  0  0  0  1  1  1  1  1  1  1  0  0  1  1  1  1  1  1
//      1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31

// 4月  001100000000000000000000000000      30天
//      0 0 1 1 0 0 0 0 0  0  0  0  0  0  0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
//      1 2 3 4 5 6 7 8 9 10

var bitMap3 = 1743814463  // 二进制bitmap对应的十进制表示
var bitMap4 = 201326592   // 二进制bitmap对应的十进制表示

func main() {
	//getSignUrl()

	//fmt.Printf("%b \n", bitMap3)
	//fmt.Printf("%b \n", bitMap4)

	//Sign("2022-04-03")
	//CalSignRewardOnMonth("2022-04-06")
	//return


	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		return
	}
	log.Println("http server exit")
}
