# Matrix Mutants

This application is a tool created to aid **Magneto** in his quest for mutants. Based on a DNA sequence this app can tell if it is a mutant or not. 

There are two endpoints on this application:

- https://arcane-dusk-93810.herokuapp.com/stats [GET]
- https://arcane-dusk-93810.herokuapp.com/mutants [POST]

## Application Usage 
Where you will receive as parameter an array of Strings that represent each row of a table of (NxN) with the DNA sequence. The letters of the Strings can only be: (A, T, C, G), which represents each nitrogenous DNA base.

![Table Image](https://github.com/viniciusfeitosa/matrix_mutantes/images/table.png)

You will know if a human is mutant, if you find **more than one sequence of four letters
equal, obliquely, horizontally or vertically**.


### Endpoint /stats

This endpoint returns a **JSON** with the statistics about the DNA validations
https://arcane-dusk-93810.herokuapp.com/stats [GET]
```javascript
{“count_mutant_dna”:40, “count_human_dna”:100: “ratio”:0.4}
```
If you try to access the endpoint ***/stats*** using another HTTP verb different then GET, will be returned the error code 405

### Endpoint /mutants

This endpoint will validate a DNA sequence to check if there is a mutant or not
https://arcane-dusk-93810.herokuapp.com/mutants [POST]
The POST body must be a JSON similar the bellow example
```javascript
{ 
	"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"] 
}
```

 - If the DNA sequence validate is **positive** to mutants will be returned the HTTP code 200. 
  - If the DNA sequence validate is **negative** to mutants will be returned the HTTP code 403. 
  - If you try to access the endpoint ***/mutants*** using another HTTP verb different then POST, will be returned the error code 405. 
  - If the DNA sequence is invalid, lettern different than ATCG or has a invalid size, will be returned the error code 500.

## About the Architecture

Was used for this app Go (golang) and Redis in two different ways, as DB, and as the queue. By the cost, DB and Queue are the same Redis instance.

Let's understand the POST flow:

 1. The user send a request with a DNA sequence in the body
 2. The handler get the data and send to the checker
 3. The checker user 4 goroutines to validade if there is a mutant on the DNA sequence. Asynchronously put the info on the queue
 4. Workers (10 goroutines) consume the data from the queue and consolidate the info on the DB

Now, let's understand the GET flow:
1. The user send a request to the aplication. 
2. The handler send the data to the stats consumer.
3. The consumer get the data from the DB. 
4. The consumer send the statistics data to the user

The bellow diagram could ilustrate the general application work: 

![Diagram Image](https://github.com/viniciusfeitosa/matrix_mutantes/images/mutants.png)

With this approach the database is not so stressed and has more freedom to consolidate the information received.

## Running the tests

To run the application test is just execute the following command:
```bash
go test ./... -v -coverprofile=. -timeout 30s
``` 

