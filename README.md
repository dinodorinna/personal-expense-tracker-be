# Personal Expense Tracker Backend

Note: This is my second Golang project, created as part of my learning journey.

A REST API backend service for tracking personal expenses built with Go, Gin framework, and SQLite database.

## Features

- **Transaction Management**: Create, Retrieve, Update and Delete expense transactions
- **Date-based Filtering**: Query transactions by specific dates
- **Pagination**: Paginated results for transaction listing

## TODO / Future Enhancements

### CRUD Operations

- [x] Implement GET /transactions
- [x] Implement request body parsing for POST /transaction endpoint
- [x] Implement DELETE /delete/:id endpoint (delete transaction)
- [ ] Add GET /transaction/:id endpoint (get single transaction)

### Reporting Features

- [ ] Get total monthly report (GET /reports/monthly/:year/:month)
- [ ] Get category report by month (GET /reports/categories/:year/:month)

### Additional Features

- [ ] Add logging
- [ ] Add unit tests
- [ ] Improvement middleware
