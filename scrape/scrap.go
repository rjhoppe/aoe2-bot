package scrape

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func ScrapeStats() error {
	err := playwright.Install()
	if err != nil {
		return fmt.Errorf("could not install playwright: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		return fmt.Errorf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		return fmt.Errorf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		return fmt.Errorf("could not create page: %v", err)
	}

	// Go to the page
	_, err = page.Goto("https://aoestats.io/civs/")
	if err != nil {
		return fmt.Errorf("could not go to page: %v", err)
	}

	// get civs from title attribute of imgs
	imgLocator := page.Locator("img")
	count, err := imgLocator.Count()
	if err != nil {
		return fmt.Errorf("could not count img elements: %w", err)
	}

	civs := make([]string, 0, 50)

	for i := 0; i < count; i++ {
		img := imgLocator.Nth(i)
		title, err := img.GetAttribute("title")
		if err != nil {
			fmt.Printf("warning: could not get title attribute: %v\n", err)
			continue
		}
		if title != "" {
			civs = append(civs, title)
		}
	}

	// get all the span class="text-stats-high"
	civWinRates := make([]string, 0, 50)

	getStats(".text-stats-high", page, &civWinRates)
	getStats(".text-stats-medium", page, &civWinRates)
	getStats(".text-stats-low", page, &civWinRates)

	// match the civ to the win rate
	civToWinRate := map[string]string{}
	for i, val := range civs {
		civToWinRate[val] = civWinRates[i]
	}

	// output civ to winrate to a file that can be read
	err = sendToLocalJsonFile(civToWinRate)
	if err != nil {
		return fmt.Errorf("error storing storing data to json: %w", err)
	}

	return nil
}

func getStats(className string, page playwright.Page, statsDataStruct *[]string) {
	statLocator := page.Locator(className)
	count, err := statLocator.Count()
	if err != nil {
		fmt.Printf("failed to count elements: %s", err)
	}

	for i := 0; i < count; i++ {
		el := statLocator.Nth(i)
		text, err := el.TextContent()
		if err != nil {
			fmt.Printf("warning: failed to get text content: %v\n", err)
			continue
		}
		cleanedText := strings.TrimSpace(text)
		if cleanedText != "" {
			*statsDataStruct = append(*statsDataStruct, cleanedText)
		}
	}
}

func sendToLocalJsonFile(data map[string]string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	// Write to a file
	fp := "data/leaderboard.json"
	err = os.WriteFile(fp, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	fmt.Println("✅ JSON file created successfully!")
	return nil
}
