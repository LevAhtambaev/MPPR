package repository

import (
	"car_bot/internal/app/ds"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"sort"
)

func GetResults(db *sqlx.DB, VkID int) (string, error) {
	var user ds.User
	err := db.QueryRow("SELECT * from users WHERE vk_id =$1", VkID).Scan(&user.VkID, &user.State, &user.CityDriving, &user.HighwayDriving, &user.FuelEfficient, &user.Spacious, &user.Powerful, &user.EasyToDrive, &user.Comfortable, &user.Sporty, &user.ModernDesign, &user.HighEndLuxury, &user.EnvironmentallyFriendly, &user.FWD, &user.RWD, &user.AWD, &user.RearViewCamera, &user.TirePressureMonitoring, &user.AutomaticParking, &user.NavigationSystem, &user.HeatedSeats, &user.ElectricMotor)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Row with id unknown")
		} else {
			log.Println("Couldn't find the line with the user")
		}
		log.Error(err)
		return "", err
	}
	cars := make(map[string]int, 20)
	log.Println(user)
	for _, car := range ds.Сars {
		if user.CityDriving == car.CityDriving {
			cars[car.Name]++
		}
		if user.HighwayDriving == car.HighwayDriving {
			cars[car.Name]++
		}
		if user.FuelEfficient == car.FuelEfficient {
			cars[car.Name]++
		}
		if user.Spacious == car.Spacious {
			cars[car.Name]++
		}
		if user.Powerful == car.Powerful {
			cars[car.Name]++
		}
		if user.EasyToDrive == car.EasyToDrive {
			cars[car.Name]++
		}
		if user.Comfortable == car.Comfortable {
			cars[car.Name]++
		}
		if user.Sporty == car.Sporty {
			cars[car.Name]++
		}
		if user.ModernDesign == car.ModernDesign {
			cars[car.Name]++
		}
		if user.HighEndLuxury == car.HighEndLuxury {
			cars[car.Name]++
		}
		if user.EnvironmentallyFriendly == car.EnvironmentallyFriendly {
			cars[car.Name]++
		}
		if user.FWD == car.FWD {
			cars[car.Name]++
		}
		if user.RWD == car.RWD {
			cars[car.Name]++
		}
		if user.AWD == car.AWD {
			cars[car.Name]++
		}
		if user.RearViewCamera == car.RearViewCamera {
			cars[car.Name]++
		}
		if user.TirePressureMonitoring == car.TirePressureMonitoring {
			cars[car.Name]++
		}
		if user.AutomaticParking == car.AutomaticParking {
			cars[car.Name]++
		}
		if user.NavigationSystem == car.NavigationSystem {
			cars[car.Name]++
		}
		if user.HeatedSeats == car.HeatedSeats {
			cars[car.Name]++
		}
		if user.ElectricMotor == car.ElectricMotor {
			cars[car.Name]++
		}
	}
	log.Println(cars)
	res := make([]string, 0, len(cars))

	for k := range cars {
		res = append(res, k)
	}

	sort.SliceStable(res, func(i, j int) bool {
		return cars[res[i]] > cars[res[j]]
	})

	var output string
	output += "Машина\tСовместимость\n"
	for _, v := range res {
		output += fmt.Sprintf("%s\t%d%%\n", v, cars[v]*100/20)
	}
	return output, nil
}
