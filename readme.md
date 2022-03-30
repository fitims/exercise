# Project description #

Create an API (consuming and producing JSON) which allows users to register, persist mazes they 
create in the database and get solutions for those mazes. Please read the following instructions 
carefully, not following them will reflect negatively on your application. 
You should implement the necessary logic for the following flow:

1. User registers via ```POST /user``` endpoint with the following fields:
   a. username (i.e. happyUser)
   b. password (i.e. iTk19!n)
2. User logs in via ```POST /login``` endpoint
3. The API creates a session and responds with a token

**From this point on all the mentioned endpoints below should require a valid token to be supplied**

5. User creates a maze via a POST /maze endpoint with the following fields:
   a. gridSize (size of a maze grid i.e. 10x10)
   b. walls (an array of cells which contain a wall within a given grid)
   c. entrance (the cell where the path should begin i.e. A1)
6. User sends a request to ```GET /maze/{mazeId}/solution endpoint``` with steps query parameter which can be either ```min``` or ```max```
7. The API returns an array of grid cells leading from the entrance of the maze to the exit of the maze with the following rules: 
   a. if ```steps``` parameter is ```min``` the API returns the path from the entrance to the exit with the **least number of steps possible**
   b. if ```steps``` parameter is ```max``` the API returns the path from the entrance to the exit with the **most number of steps possible**
8. User can see their created mazes by sending a request to GET /maze (the user should be able to see just their own mazes)

**Example:**

```POST /maze``` request body:

```
{
    "entrance":  "A1",
    "gridSize": "8x8",
    "walls": ["C1", "G1", "A2", "C2", "E2", "G2", "C3", "E3", "B4", "C4", "E4", "F4", "G4", "B5", "E5", "B6", "D6", "E6", "G6", "H6", "B7", "D7", "G7", "B8"]`
}
```

```GET /maze/1/solution?steps=min``` response:

```
{
    "path": ["A1", "B1", "B2", "B3", "A3", "A4", "A5", "A6", "A7", "A8"]`
}
```

**Additional notes:**

- If the maze has no solution an error should be thrown
- API needs to detect the exit point automatically
- A maze can only have one exit point (at one the edge cells of the grid), otherwise an error should be thrown