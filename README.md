# CookMe — REST API

Go REST API for the CookMe recipe & grocery ordering platform.

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.21 |
| Web Framework | [Fiber v2](https://github.com/gofiber/fiber) |
| ORM | GORM + PostgreSQL |
| Search | Elasticsearch 8 |
| Maps | Google Maps Platform (Places, Distance Matrix) |
| Config | godotenv |

## Project Structure

```
COOKME_rest-api/
├── main.go                  # Entry point — wires up services and starts server
├── server/server.go         # Route definitions
├── handlers/                # HTTP request/response layer
├── core/                    # Business logic
├── repositories/            # Database access (GORM)
├── models/                  # Struct definitions
├── ports/                   # Interfaces (dependency inversion)
└── SuperMarketPrices/       # Supermarket ingredient price data files
```

## Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- Google Maps API key (Places API + Distance Matrix API)
- Elasticsearch (included in docker-compose)

## Getting Started

### 1. Configure environment

Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=cookme
DB_PASS=cookme
DB_NAME=cookme
```

### 2. Start the database and Elasticsearch

```bash
# On WSL (use the Unix socket)
DOCKER_HOST=unix:///var/run/docker.sock docker.exe compose up -d

# On native Linux / macOS
docker compose up -d
```

The database is automatically populated from `backup.sql` on first start.

### 3. Run the API

```bash
go run main.go
```

The server starts on `http://localhost:3000`.

## API Reference

### Recipes

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/recipe/` | Get all recipes |
| `GET` | `/recipe/:id` | Get recipe by ID |
| `POST` | `/recipe/` | Create a recipe |
| `POST` | `/recipe/recipes` | Bulk create recipes |
| `DELETE` | `/recipe/:id` | Delete a recipe |
| `GET` | `/recipe/cuisines` | Filter recipes by cuisine |
| `GET` | `/recipe/keywords` | Filter recipes by keywords |
| `POST` | `/recipe/convertToMarketIngredients/:recipeID?market=AB` | Get priced ingredients for a recipe at a given market |
| `POST` | `/recipe/compareMarketPrices/:recipeID` | Find the cheapest market for a recipe |

### Users

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/user/` | Register a new user |
| `POST` | `/user/login` | Login |
| `GET` | `/user/:id` | Get user by ID |
| `GET` | `/user/email/:email` | Get user by email |
| `GET` | `/user/all` | Get all users |
| `DELETE` | `/user/:id` | Delete user |
| `GET` | `/user/profileRecipes/:userId` | Get user's recipe activity |
| `GET` | `/user/closestMarket/:userId` | Find nearest supermarket to user |
| `GET` | `/user/availableMarkets/:userId` | Find all supermarkets within 5km |
| `POST` | `/user/addFavoriteRecipe/:userId/:recipeId` | Add recipe to favourites |
| `DELETE` | `/user/removeFavoriteRecipe/:userId/:recipeId` | Remove from favourites |
| `GET` | `/user/favoriteRecipes/:userId` | List favourite recipes |
| `POST` | `/user/addReview` | Add a review |
| `PUT` | `/user/updateReview` | Update a review |
| `GET` | `/user/reviews/:userId` | Get reviews by user |

### Orders

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/order/` | Place an order |
| `GET` | `/order/:userId` | Get orders for a user |
| `GET` | `/order/all` | Get all orders (admin) |
| `DELETE` | `/order/:recipeId/:userId` | Cancel an order |

### Search

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/search/:keyword` | Full-text search via Elasticsearch |
| `GET` | `/search/` | Get all searchable keywords |

## Notes

- The Elasticsearch service is required at startup. If it is unavailable the API will fail to start.
- Supermarket location lookups require a Google Maps API key with **Places API** and **Distance Matrix API** enabled, set directly in `core/userCore.go`.
- The `backup.sql` init script only runs once — on a fresh volume. To re-seed, run `docker compose down -v` first.
