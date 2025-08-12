# 📈 Yahoo Finance Stock Scraper (Go)

A lightweight Go application that scrapes real-time stock data from [Yahoo Finance](https://finance.yahoo.com/) for a predefined list of ticker symbols.  
It outputs the results to the console and saves them in a CSV file for later analysis.

---

## 🚀 Features
- Scrapes **Company Name**, **Current Price**, and **Price Change**.
- Preconfigured ticker list (can be easily modified).
- Saves results to `stocks.csv`.
- Lightweight — built with [Colly](https://github.com/gocolly/colly).

---

## 🛠 Requirements
- **Go** 1.18 or later
- **Git**
- Internet connection

---

## 📦 Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/stock-scraper-go.git
   cd stock-scraper-go
   ```

2. Install dependencies:
   ```bash
   go mod init stock-scraper
   go get github.com/gocolly/colly
   ```

---

## ▶ Usage
Run the scraper:
```bash
go run main.go
```

After running:
- The console will display each stock’s details.
- A file named `stocks.csv` will be generated with the results.

---

## 📊 Example Output

**Console:**
```
Visiting https://finance.yahoo.com/quote/MSFT
Company: Microsoft Corporation, Price: 327.12, Change: +2.15
...
Scraping completed successfully.
Total stocks scraped: 12
```

**CSV (`stocks.csv`):**
| Company                 | Price   | Change |
|-------------------------|---------|--------|
| Microsoft Corporation   | 327.12  | +2.15  |
| Apple Inc.              | 172.05  | -1.23  |

---

## ⚙ Customization
To change the list of tickers, modify the `ticker` slice in `main.go`:
```go
ticker := []string{"MSFT", "AAPL", "GOOGL"}
```

If Yahoo Finance changes its HTML structure, update the selectors in:
```go
c.OnHTML("#main-content-wrapper", func(e *colly.HTMLElement) { ... })
```

---

## ⚠ Notes
- Avoid sending too many requests too quickly — you may get blocked.
- The scraper relies on Yahoo Finance's HTML layout, which may change.

---

## 📄 License
MIT License
