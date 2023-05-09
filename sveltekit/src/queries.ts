import HttpClient from './REST.js'
//Fetch service for appointment

let base_url = 'http://localhost:8080'
//extends the Rest class 
class ApiClient extends HttpClient {

    constructor() {
        super({
            baseURL: `${base_url}`,
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'Access-Control-Allow-Origin': '*',
                'Access-Control-Allow-Headers': '*',
            }
        })
    }
// methods
		/**
		 * get getAlbums will return all albums as a list.
		 * 
		 */
    get getAlbums() {

        return {
            get: () => this.get("/albums")
        }
    }
    /**
		 * get getAlbums will return all albums as a list.
		 * 
		 */
    get getDefault() {

        return {
            get: () => this.get("/")
        }
    }
    

		/**
		 * get getAlbum will take the albumId and return that album.
		 * 
		 * ## Parameters
		 * 
		 * - albumId - The ID of the album we want to view.
		 */
		get getAlbum(){
			return {
				get: (albumId: Number) => this.get(`/album/${albumId}`)
			}
		}

		/**
		 * get addAlbum will take the JSON stringified album and add it to the backend database.
		 * 
		 * 
		 * ## Parameters
		 * 
		 * - album - The album we want to add..
		 */
		get addAlbum(){
			return {
				post: (album: Object) => this.post(`/albums`, album)
			}
		}
}
export default ApiClient

