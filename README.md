# game_sales_api

## Endpoints

### /games/all

Returns all data in paginated form. If no parameters are provided this will simply return the top 50 results.

#### URL Parameters

- page - The page of results to return
- resultsPerPage - Number of results per page (up to 50)

### /games/{rank}

Returns the game at a specific ranking number

### /games/search

Searches the data based upon specified parameters.

#### Search Query URL Parameters

- **name** - The name of a specific game (will do approximate match)
- **platform** - A specifit platform (ping /platforms for options)
- **genre** - A specific genre (ping /genres for options)
- **publisher** - A specific publisher (ping /publishers for options)
- **minSales** - The minimum number of global sales a title should have achieved in millions, shorted to decimal form (eg: 3.01 for 3,100,000)
- **maxSales** - The maximum number of global sales a title should have achieved in millions, shorted to decimal form.
- **orderby**
  - **"asc"** - Sales Ascending
  - **"desc"** - Sales Descending

### /platforms

Returns a list of valid platforms. Includes total global sales for each platform in millions.

### /genres

Returns a list of valid genres. Includes total global sales for each genre in millions.

### /publishers

Returns a list of valid publishers. Includes total global sales for each publisher in millions.
