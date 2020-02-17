package validation

import (
	"Backend_task_RF/DB"
	"errors"
	"strconv"
)

func ValidateLimit(limit string) (resultLimit string, err error) {
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return "", errors.New("Задано некорректное значение")
		}
		if limitInt < 0 {
			return "", errors.New("Значение не может быть отрицательным")
		} else if limitInt == 0 {
			return "", errors.New("Значение не может быть нулевым")
		}
		
		numberOfRecords, err := gettingNumberOfRecords()
		if err != nil{
			return "", err	
			}

		if limitInt > numberOfRecords {
		   return "limit" + " " + strconv.Itoa(numberOfRecords), nil
		} else {
			return "limit" + " " + limit, nil
		}
		return "", nil
	}
    return "", nil
}

func ValidateOffset(offset string) (resultOffset string, err error) {
	if offset != "" {
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			return "", errors.New("Задано некорректное значение")
		}
		
		if offsetInt < 0 {
		   return "", errors.New("Значение не может быть отрицательным")
		}

		numberOfRecords, err := gettingNumberOfRecords()
		if err != nil{
		   return "", err	
		}
		
		if offsetInt > numberOfRecords {
		   return "offset" + " " + strconv.Itoa((numberOfRecords - 1)), nil
		} else {
		    return "offset" + " " + offset, nil
		}
		return "", nil
	}
	return "", nil
}

func gettingNumberOfRecords()(numbeOfRecords int, err error){
	
	db, err := DB.Open()
	if err != nil {
		return 0, err
	}
	
	query := "SELECT count(*) FROM the_users tu "
	rows, err := db.Query(query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var numberOfRecords int
	for rows.Next() {
		if err = rows.Scan(&numberOfRecords); err != nil {
			return 0, err
		}
	}
	
	if err = rows.Err(); err != nil {
		return 0, err
	}
	return numberOfRecords, nil
}