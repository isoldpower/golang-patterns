package src

import (
	"bufio"
	"fmt"
	"go-patterns/internal/util"
	"go-patterns/src/pipeline"
	"os"
	"strconv"
	"strings"
)

func readPatternChoice(patternsMap map[string]util.Command) (util.Command, bool) {
	reader := bufio.NewReader(os.Stdin)
	choicesArray := make([]util.Command, len(patternsMap))
	iterator := 0

	for key := range patternsMap {
		fmt.Printf("  %d) %s\n", iterator+1, key)
		choicesArray[iterator] = patternsMap[key]
		iterator++
	}

	finalChoice, readError := reader.ReadString('\n')
	trimmedChoice := strings.TrimSpace(finalChoice)

	parsedChoice, parseError := strconv.Atoi(trimmedChoice)
	if readError != nil || parseError != nil {
		return nil, false
	}

	return choicesArray[parsedChoice-1], true
}

func ExecuteProgram() {
	var patternsMap = map[string]util.Command{
		"pipeline": pipeline.NewCommand(),
	}

	fmt.Println("Which pattern do you want to execute?")

	if selectedPattern, ok := readPatternChoice(patternsMap); ok {
		selectedPattern.Execute()
	} else {
		fmt.Println("Pattern not found. Exiting...")
	}
}
