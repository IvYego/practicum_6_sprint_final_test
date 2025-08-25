package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoConvert(input string) (string, error) {

	if input == "" {
		return "", errors.New("input is empty")
	}

	if strings.ContainsAny(input, ".-") {
		return morse.ToText(input), nil
	}

	if strings.ContainsAny(input, "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя") {
		return morse.ToMorse(input), nil
	}

	return input, nil

}
