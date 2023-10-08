package ds

type User struct {
	VkID                    int    `db:"vk_id"`
	State                   string `db:"state"`
	CityDriving             bool   `db:"city_driving"`
	HighwayDriving          bool   `db:"highway_driving"`
	FuelEfficient           bool   `db:"fuel_efficient"`
	Spacious                bool   `db:"spacious"`
	Powerful                bool   `db:"powerful"`
	EasyToDrive             bool   `db:"easy_to_drive"`
	Comfortable             bool   `db:"comfortable"`
	Sporty                  bool   `db:"sporty"`
	ModernDesign            bool   `db:"modern_design"`
	HighEndLuxury           bool   `db:"high_end_luxury"`
	EnvironmentallyFriendly bool   `db:"environmentally_friendly"`
	FWD                     bool   `db:"fwd"`
	RWD                     bool   `db:"rwd"`
	AWD                     bool   `db:"awd"`
	RearViewCamera          bool   `db:"rear_view_camera"`
	TirePressureMonitoring  bool   `db:"tire_pressure_monitoring"`
	AutomaticParking        bool   `db:"automatic_parking"`
	NavigationSystem        bool   `db:"navigation_system"`
	HeatedSeats             bool   `db:"heated_seats"`
	ElectricMotor           bool   `db:"electric_motor"`
}
