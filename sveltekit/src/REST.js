
// Httpclient as class
class HttpClient {
    constructor(options = {}) {
        // @ts-ignore
        this._baseURL = options.baseURL || ""
        // @ts-ignore
        this._headers = options.headers || {}
    }
    // passing parameters to fetch, like header URL and Endpoints
    // @ts-ignore
    async _fetchAPI(endpoint, options = {}) {
        const res = await fetch(this._baseURL + endpoint, {
            ...options,
            headers: this._headers
        })

        //throws error if response is not OK
        if(!res.ok) throw new Error(res.statusText)
        
        // @ts-ignore
        if(options.parseResponse !== false && res.status !== 204)
            return res

        return undefined
    }
    // setting Header for fetch
    /**
	 * @param {string | number} key
	 * @param {any} value
	 */
    setHeader(key, value) {
        this._headers[key] = value
        return this
    }

    // @ts-ignore
    getHeader(key) {
        return this._headers[key]
    }

    // We can also set Authorization and bearerAuth token here if future to add security to our API endpoints


    // Methods for fetch
    // @ts-ignore
    get(endpoint, options = {}) {
        return this._fetchAPI(endpoint, {
            ...options,
            method: "GET"
        })
    }
    // @ts-ignore
    post(endpoint, body, options={}) {
        return this._fetchAPI(endpoint, {
            ...options,
						body: JSON.stringify(body),
            method: "POST"
        })
    }
    // @ts-ignore
    put(endpoint, body, options={}) {
        return this._fetchAPI(endpoint, {
            ...options,
            body: body ? JSON.stringify(body) : undefined,
            method: "PUT"
        })
    }
    // @ts-ignore
    patch(endpoint, operations, options={}) {
        return this._fetchAPI(endpoint, {
            parseResponse: false,
            ...options,
            body: JSON.stringify(operations),
            method: "PATCH"
        })
    }
    // @ts-ignore
    delete(endpoint, options={}) {
        return this._fetchAPI(endpoint, {
            parseResponse: false,
            ...options,
            method: "DELETE"
        })
    }
}

export default HttpClient