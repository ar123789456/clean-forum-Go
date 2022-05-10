import Models from './models.js'
import Posts from "./pages/allposts.js"
import Signup from "./pages/signup.js"
import Signin from "./pages/signin.js"

export default {
    async postsRoute() {
        const posts = await Models.posts()
        console.log(posts);
        Posts.setData(posts)
        Posts.render()
    },
    async signupRoute() {
        // const posts = await Models.signup()
        let data = {}

        Signup.setData(data)
        Signup.render()
        document.forms.signupform.onsubmit = function() {
            var message = {}
            let n = 3
            if (this.name.value != "") {
                n--
                message.name = this.name.value
            }
            if (this.email.value != "") {
                n--
                message.email = this.email.value
            }
            if (this.password.value != "") {
                n--
                message.password = this.password.value
            }
            if (n == 0) {
                let status = Models.signup(message)
                alert(status)
            }
            console.log(message)
            return false;
        };

    },
    async signinRoute() {
        // const posts = await Models.signup()
        let data = {}

        Signin.setData(data)
        Signin.render()
        document.forms.signinform.onsubmit = function() {
            var message = {}
            let n = 2
            if (this.name.value != "") {
                n--
                message.name = this.name.value
            }
            if (this.password.value != "") {
                n--
                message.password = this.password.value
            }
            if (n == 0) {
                Models.signin(message)

            }
            // console.log(message)
            return false;
        };

    },
}