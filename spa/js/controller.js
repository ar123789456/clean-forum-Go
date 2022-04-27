import Models from './models.js'
import Posts from "./pages/allposts.js"
import Signup from "./pages/signup.js"

export default {
    async postsRoute() {
        const posts = await Models.posts()
        console.log(posts);
        Posts.setData(posts)
        Posts.render()
    },
    async signupRoute() {

        Signup.setData()
        Signup.render()
    },
    async signupFormRoute() {
        const posts = await Models.signup()
        console.log(posts);
        Posts.setData(posts)
        Posts.render()
    }

}