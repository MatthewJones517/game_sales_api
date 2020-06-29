# game_sales_api

## Endpoints

### /all

Returns all data in paginated form. If no parameters are provided this will simply return the top 50 results. 

#### Query Parameters

- page - The page of results to return
- resultsPerPage - Number of results per page (up to 50)
- orderby
  - "asc" - Sales Ascending
  - "desc" - Sales Descending

### /search

Searches the data based upon specified parameters.

#### Query Parameters

- **name** - The name of a specific game (will do approximate match)
- **platform** - A specifit platform (ping /platforms for options)
- **genre** - A specific genre (ping /genres for options)
- **publisher** - A specific publisher (ping /publishers for options)
- **minSales** - The minimum number of global sales a title should have achieved in millions, shorted to decimal form (eg: 3.01 for 3,100,000)
- **maxSales** - The maximum number of global sales a title should have achieved in millions, shorted to decimal form (eg: 3.01 for 3,100,000)
- **orderby**
  - **"asc"** - Sales Ascending
  - **"desc"** - Sales Descending

### /platforms

Returns a list of valid platforms. Includes total global sales for each platform.

### /genres

Returns a list of valid genres. Includes total global sales for each genre.

### /publishers

Returns a list of valid publishers. Includes total global sales for each publisher. 

### /rank/{number}

Returns the game at a specific rank.