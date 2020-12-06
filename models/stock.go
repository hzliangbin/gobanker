package models

type Stock struct {
	Code string
	Type string
	Name string
	Price float64
	YesterdayPrice float64
	Fluctuate float64
	TodayMax float64
	TodayMin float64
	PriceDate int32
	industry string
	totalValue float64
	pb float64
	roe float64
	bvps  float64
	//市盈率 动态
	pes  float64
	//市盈率 动态
	ped float64
	//市盈率
	pettm float64
	//high52w
	high52w float64
	//52周最低
	low52w float64
	/**
	 * 静态分红日期
	 */
	dividendDate string
	dividend float64
	/**
	 * 静态分红更新时间
	 */
	dividendUpdateDay int32
	/**
	 * 实时股息率
	 */
	dy float64
	/**
	 * 5年平均股息
	 */
	fiveYearDy float64
	/**
	 * 3年平均股息
	 */
	threeYearDy float64
	/**
	 * 5年平均Roe
	 */
	fiveYearRoe float64
	/**
	 * 3年平均Roe
	 */
	threeYearRoe float64
	/**
	 * 总营业收入
	 */
	totalIncome float64
	/**
	 * 同期对比总营业收入
	 */
	incomeDiff float64
	/**
	 * 净利润
	 */
	totalProfits float64
	/**
	 * 毛利率
	 */
	mll float64
	/**
	 * 同期对比净利润
	 */
	profitsDiff float64
	/**
	 * 报告期
	 */
	report string
	/**
	 * 股票分类
	 */
	stype string
}
