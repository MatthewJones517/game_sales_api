# Game Sales API

## What Is This?

I'm a firm believer that the best way to learn a language is by programming in it. Go is no different. I started by reading through the docs to get a high level idea of how things worked, then decided to write this test project.

I found a [dataset on Kaggle](https://www.kaggle.com/gregorut/videogamesales) that contained a list of over 16k video games that sold at least 100,000 copies. I then wrote this API that allows you to access that data in some interesting ways.

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
- **platform** - A specific platform (ping /platforms for options)
- **genre** - A specific genre (ping /genres for options)
- **publisher** - A specific publisher (ping /publishers for options)
- **minSales** - The minimum number of global sales a title should have achieved in millions, shorted to decimal form (eg: 3.01 for 3,100,000)
- **maxSales** - The maximum number of global sales a title should have achieved in millions, shorted to decimal form.

### /platforms

Returns a list of valid platforms. Includes total global sales for each platform in millions.

### /genres

Returns a list of valid genres. Includes total global sales for each genre in millions.

### /publishers

Returns a list of valid publishers. Includes total global sales for each publisher in millions.
