# Initial Project Requirements

## Minimally Viable Product
* Social Login
  - Ability to login to the application via oauth2 flow with Facebook, Google
  - Use something off the shelf if possible, Auth0, StormPath, whatever
* Ability to connect to user's pinterest account
  - Authorize to browse boards
  - Select boards to pull recipes from
* Ability to break a recipe into ingredients
* Ability to categorize ingredients as "pantry items" or "single use"
  - Should be able to fully maintain information around an ingredient
  - Track type (produce, meat, dairy, etc)
  - Track durability - fresh veg only lasts X days
* Ability to add all ingredients from a recipe to grocery list
* Ability to exclude / include pantry items on the grocery list
* Ability to remove individual items from the grocery list
* Ability to remove an entire recipe of items from the grocery list
* Ability to add ad-hoc items to a grocery list
* Ability to mark items as purchased on the grocery list

## Nice to haves / Phase 2
* Ability to source durable grocery items from amazon / walmart / etc
* Ability to integrate with Alexa, Google Home, etc
* Ability to generate a meal plan as a random selection of recipes
* Ability to apply weights to meals so that they are selected more frequently based on a number of factors
* Ability to optimize pantry utilization e.g. select recipes that favor items that already exist that were purchased in bulk
* Ability to optimize produce utilization e.g. select two recipes that both require half an onion or a fraction of a bundle of cilantro - organize the meal plan so that no produce goes bad