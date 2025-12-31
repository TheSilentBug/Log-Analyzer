package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// مرحله 1: ورود فایل لاگ
func inputLogFile() string {
	// نام فایل لاگ را از ورودی کاربر می‌گیریم
	var filename string
	fmt.Print("Enter the log file name: ")
	fmt.Scanln(&filename)
	return filename
}

// مرحله 2: خواندن فایل لاگ
func readLogFile(filename string) (string, error) {
	// فایل لاگ را می‌خوانیم
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// مرحله 3: تجزیه داده‌ها
func parseData(data string) []string {
	// داده‌ها را به خطوط تقسیم می‌کنیم
	lines := strings.Split(data, "\n")
	return lines
}

// مرحله 4: شناسایی الگوها (مثل خطاها)
func identifyPatterns(lines []string) []string {
	var errors []string
	for _, line := range lines {
		if strings.Contains(line, "ERROR") { // به عنوان مثال خطاهای "ERROR" را شناسایی می‌کنیم
			errors = append(errors, line)
		}
	}
	return errors
}

// مرحله 5: پردازش داده‌ها
func processData(errors []string) {
	// در اینجا می‌توانیم تحلیل‌هایی مثل تعداد خطاها را انجام دهیم
	fmt.Printf("Total errors found: %d\n", len(errors))
}

// مرحله 6: تولید گزارش
func generateReport(errors []string) {
	// گزارش تولید می‌شود
	fmt.Println("Generating report...")
	for _, err := range errors {
		fmt.Println(err)
	}
}

// مرحله 7: خروجی (نمایش یا ذخیره نتایج)
func outputResults() {
	// در اینجا می‌توانیم گزارش را نمایش دهیم یا در فایل ذخیره کنیم
	fmt.Println("Output results displayed.")
}

// مرحله 8: پایان
func main() {
	// اجرای گام‌ها به ترتیب
	fmt.Println("Log Analyzer Started")

	// ورود فایل لاگ
	filename := inputLogFile()

	// خواندن فایل لاگ
	data, err := readLogFile(filename)
	if err != nil {
		fmt.Println("Error reading log file:", err)
		os.Exit(1)
	}

	// تجزیه داده‌ها
	lines := parseData(data)

	// شناسایی الگوها
	errors := identifyPatterns(lines)

	// پردازش داده‌ها
	processData(errors)

	// تولید گزارش
	generateReport(errors)

	// خروجی نتایج
	outputResults()

	fmt.Println("Log Analyzer Completed")
}
