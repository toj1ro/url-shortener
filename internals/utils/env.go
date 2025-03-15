package utils

import (
	"bufio"
	"os"
	"strings"
)

// Загружает переменные из .env файла в окружение
func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err // Файл не найден
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Пропускаем пустые строки и комментарии
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Разделяем ключ и значение
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Некорректный формат
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Удаляем кавычки (если есть)
		if (strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`)) ||
			(strings.HasPrefix(value, `'`) && strings.HasSuffix(value, `'`)) {
			value = value[1 : len(value)-1]
		}

		// Устанавливаем переменную, если она ещё не задана
		if os.Getenv(key) == "" {
			err := os.Setenv(key, value)
			if err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}
