package controllers

import (
	"github.com/revel/revel"
	"github.com/Nicolas-Constanty/Dnai.ML.API/app/models"
	"fmt"
	"os"
	"bufio"
	"path/filepath"
	"strings"
	"strconv"
	"time"
)

type App struct {
	*revel.Controller
}
func (c App) Maps() revel.Result{
	return c.Render()
}
func readLine(path string) []string {
	var dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if (err != nil) {
		fmt.Println(err)
	}
	inFile, _ := os.Open(dir + "/" + path)
	defer inFile.Close()
	var lines []string
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func (c App) InitialValue() revel.Result {
	data := models.Building{}
	data.Name = "Building"
	lines := readLine("resources/datas2d")
	for _, line := range lines {
		s := strings.Split(line, " ")
		x, err := strconv.ParseFloat(s[0], 64)
		if (err != nil) {
			fmt.Println(err)
		}
		y, err := strconv.ParseFloat(s[1], 64)
		if (err != nil) {
			fmt.Println(err)
		}
		pos := models.Vector2{ Lat: x, Lng: y }
		def := models.Deformation2{ Position:pos, Deformation: 0, Time: time.Time{} }
		data.Values = append(data.Values, def)
	}
	return c.RenderJSON(data)
}

func (c App) Index() revel.Result {
	return c.Render()

}
