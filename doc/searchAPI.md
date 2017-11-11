# Search API
## REST APi


**Get search results by keyword**
----
Returns the merged keyword result for search term in specified platforms


* **URL**

  /search/libs/:keyword[?queryParams]

* **Method:**

  `GET`
  
*  **URL Params**


	1. Path and/or Resource
   `keyword=[string]`

	2. Query Params
	    * Requireds:
	    None

	    * Optionals:
	    
	    	1. Ordered sub-list of excluded platforms from search
	    		    ```excluded=(string,)*```
	    	2. Ordered sub-list of version ids 
	    		    ```versions=(string,)*```
	    		    *if not present, latest version is returned*
	    		    *Note: version ids could have regular expression using * char, for instance, 4.1.* returns all versions under 4.1. prefix*

	    	3. Branch Name List: (for github)
      	    		    ```branches=(string,)*```
      	    		    *if not present, master branch is returned*

	    	4. Commit IDs List: (for github)
			    ```commitIDs=(string,)*```
			    *if not present, latest version is returned*
	    	

	

*  **Header Params**
 
	1.  content-type:
		* application/json
			*mandatory*
		 (only application/json is currently supported)
	2. Authorization:
		* Bearer JWTTokenValue
			*mandatory*
		

* **Data Params**

```javascript
{
		id : '12dsljkwq9233809', 
		app: 'search.libs',
		version: 'v0.1',
		startId: '347289fsdkjhdf',
		direction: 'up',
		metaData: {
		rowCount: 10},
		
		
}
```		
* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
```javascript
{
		id : '12dsljkwq9233809', 
		app: 'search.libs',
		version: 'v0.1',
		metaData: {
		rowCount: 10 },
		resultSet: [
		{
			id:'1',
			repoType:'maven',
			params: {
				groupId: '',
				artifactId: ''
				},
			imgURI:'',
			description: ' '} ,
		
			{
			id:'2',
			repoType:'maven',
			params: {
				groupId: '',
				artifactId: ''
				},
			imgURI:'',
			description: ' '} ]
}
```
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** 
```javascript
    {
	id : '12dsljkwq9233809', 
	statusCode: '404',
	app: 'search.libs',
	version: 'v0.1',
	errorCode: '01',
	error : 'User doesn't exist', 
	errorPage: "https://causalera.io/errors/01',
	supportErrorCode: '111201'
	}
```

  * **Code:** 401 UNAUTHORIZED
    **Content:**
```javascript
    {
    	id : '12dsljkwq9233809', 
	statusCode: '401',
	app: 'search.libs',
	version: 'v0.1',
	errorCode: '01',
	error : 'You are unauthorized to make this request.', 
	errorPage: 'https://causalera.io/errors/01',
	supportErrorCode: '111201'
	}
```


  * **Code:** 500 Internal Server Error
    **Content:**
```javascript
    {
    	id : '12dsljkwq9233809', 
	statusCode: '500',
	app: 'search.libs',
	version: 'v0.1',
	errorCode: '01',
	error : 'Internal Server Error.',
	errorPage: 'https://causalera.io/errors/01',
	supportErrorCode: '111201'
	}
```



* **Sample Call:**

```curl
curl --request GET \
  --url 'https://localhost/v0_1/search/libs/guice?excluded=github%2Ccausalera' \
  --header 'authorization: Bearer aaaaaaaaaaaaa.bbbbbbbbbbbbbb.cccccccccccccc' \
  --header 'content-type: application/json' \
  --data '{
  
		id : '\''12dsljkwq9233809'\'', 
		app: '\''search.libs'\'',
		version: '\''v0.1'\'',
		startId: '\''347289fsdkjhdf'\'',
		direction: '\''up'\'',
		metaData: {
		rowCount: 10},
}'
```

