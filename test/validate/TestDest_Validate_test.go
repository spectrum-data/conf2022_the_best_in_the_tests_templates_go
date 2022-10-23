package validate

import (
	"conf2022_the_best_in_the_tests_templates_go/input"
	"os"
	"strings"
	"testing"
)

func CreateTempFileWithContext(context string) string {
	file, err := os.CreateTemp(os.TempDir(), "test_descriptions_*.csv")

	if err != nil {
		panic(err)
	}

	file.WriteString(context)
	file.Sync()
	file.Close()

	return file.Name()
}

func Test_EmptyFile_invalid(t *testing.T) {
	filePath := CreateTempFileWithContext("")

	defer os.Remove(filePath)

	validation := input.Validate(filePath)

	if validation.IsValid {
		t.Error("Пустой файл - не валидный файл")
	}
}

func Test_OnlyHeader_valid(t *testing.T) {
	filePath := CreateTempFileWithContext(input.DEFAULT_HEADER)

	defer os.Remove(filePath)

	validation := input.Validate(filePath)

	if !validation.IsValid {
		t.Error("Файл, содержащий один только заголовок, должен быть валиден")
	}
}

func Test_IncorrectNumberDelimiters_invalid(t *testing.T) {
	incorrectLine := strings.Repeat(string(input.DEFAULT_COLUMN_DELIMITER), input.DEFAULT_FIELDS_COUNT+1)

	filePath := CreateTempFileWithContext(
		strings.Join(
			[]string{input.DEFAULT_HEADER, incorrectLine},
			"\n"),
	)

	defer os.Remove(filePath)

	validation := input.Validate(filePath)

	if validation.IsValid {
		t.Error("Файл, содержащий строку с неверным кол-вом разделителей, должен быть не валидным")
	}
}

func Test_CorrectNumberDelimiters_valid(t *testing.T) {
	correctLine := strings.Join([]string{"harisov", "1", "someTestCase", "false", "someComment"}, string(input.DEFAULT_COLUMN_DELIMITER))

	var filePath = CreateTempFileWithContext(
		strings.Join(
			[]string{input.DEFAULT_HEADER, correctLine},
			"\n"),
	)

	defer os.Remove(filePath)

	validation := input.Validate(filePath)

	if !validation.IsValid {
		t.Error("Файл, содержащий хеден и валидную строку, должен быть валидным")
	}
}

func Test_IncorrectNumberFormat_invalid(t *testing.T) {
	incorrectNumber := "anything, but not the number"
	correctLine := strings.Join([]string{"harisov", incorrectNumber, "someTestCase", "false", "someComment"}, string(input.DEFAULT_COLUMN_DELIMITER))

	var filePath = CreateTempFileWithContext(
		strings.Join(
			[]string{input.DEFAULT_HEADER, correctLine},
			"\n"),
	)

	defer os.Remove(filePath)

	validation := input.Validate(filePath)

	if validation.IsValid {
		t.Error("Файл, содержащий строку с некоторректным номером, должен быть не валидным")
	}
}

func Test_DuplicateTestId_invalid(t *testing.T) {
	firstLine := strings.Join([]string{"harisov", "1", "someTestCase", "false", "someComment"}, string(input.DEFAULT_COLUMN_DELIMITER))
	secondLine := strings.Join([]string{"harisov", "1", "anotherSomeTestCase", "true", "anotherSomeComment"}, string(input.DEFAULT_COLUMN_DELIMITER))

	var filePath = CreateTempFileWithContext(
		strings.Join(
			[]string{input.DEFAULT_HEADER, firstLine, secondLine},
			"\n"),
	)

	defer os.Remove(filePath)

	validation := input.Validate(filePath)

	if validation.IsValid {
		t.Error("Файл, содержащий совпадающие связки автора теста и идентификатору, должен быть не валидным")
	}
}
