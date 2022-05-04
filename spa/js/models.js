export default {
    //post/////////////////////////////////////////////////
    //get
    async posts() {
        let response = await fetch("http://localhost:8080/");

        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            let json = await response.json();
            return json
        } else {
            return "Ошибка HTTP: " + response.status
        }
    },
    async singlePost(id) {
        let response = await fetch("http://localhost:8080/post/" + id);

        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            let json = await response.json();
            return json
        } else {
            return "Ошибка HTTP: " + response.status
        }
    },
    async userPosts(id) {
        let response = await fetch("http://localhost:8080/post/user/" + id);

        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            let json = await response.json();
            return json
        } else {
            return "Ошибка HTTP: " + response.status
        }
    },
    async userLikedPosts(id) {
        let response = await fetch("http://localhost:8080/post/liked/user/" + id);

        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            let json = await response.json();
            return json
        } else {
            return "Ошибка HTTP: " + response.status
        }
    },
    //create
    async postCreate(data) {
        const response = await fetch("http://localhost:8080/post/new", {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            mode: 'cors', // no-cors, *cors, same-origin
            cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
            credentials: 'same-origin', // include, *same-origin, omit
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            redirect: 'follow', // manual, *follow, error
            referrerPolicy: 'no-referrer', // no-referrer, *client
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            return response.status
        } else {
            alert(response.status)
            return "Ошибка HTTP: " + response.status
        }
    },
    //sign/////////////////////////////////////////////////
    async signup(data) {
        const response = await fetch("http://localhost:8080/signup", {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            mode: 'cors', // no-cors, *cors, same-origin
            cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
            credentials: 'same-origin', // include, *same-origin, omit
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            redirect: 'follow', // manual, *follow, error
            referrerPolicy: 'no-referrer', // no-referrer, *client
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            return response.status
        } else {
            alert(response.status)
            return "Ошибка HTTP: " + response.status
        }
    },
    async signin(data) {
        const response = await fetch("http://localhost:8080/signin", {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            crossDomain: true,
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            console.log(response.headers)
            alert()

            return response.status
        } else {
            alert(response.status)
            return response.status
        }
    },
    //categiries///////////////////////////////////////////
    //getall
    async categories() {
        let response = await fetch("http://localhost:8080/category/create");

        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            let json = await response.json();
            return json
        } else {
            return "Ошибка HTTP: " + response.status
        }
    },
    //create
    async categoryCreate(data) {
        const response = await fetch("http://localhost:8080/categories", {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            crossDomain: true,
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            console.log(response.headers)
            alert()

            return response.status
        } else {
            alert(response.status)
            return response.status
        }
    },
    //categiries///////////////////////////////////////////
    //getall
    async tags() {
        let response = await fetch("http://localhost:8080//tags");

        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            let json = await response.json();
            return json
        } else {
            return "Ошибка HTTP: " + response.status
        }
    },
    //create
    async tagCreate(data) {
        const response = await fetch("http://localhost:8080/tag/create", {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            crossDomain: true,
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            console.log(response.headers)
            alert()

            return response.status
        } else {
            alert(response.status)
            return response.status
        }
    },
    //comments////////////////////////////////////////////
    //get post comment
    async commentGetById(id) {
        let response = await fetch("http://localhost:8080/comment/" + id);

        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            let json = await response.json();
            return json
        } else {
            return "Ошибка HTTP: " + response.status
        }
    },
    //create
    async commentCreate(data) {
        const response = await fetch("http://localhost:8080/comment/new", {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            crossDomain: true,
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            console.log(response.headers)
            alert()

            return response.status
        } else {
            alert(response.status)
            return response.status
        }
    },
    //delete
    async commentDelete(data) {
        const response = await fetch("http://localhost:8080/comment/delete", {
            method: 'DELETE', // *GET, POST, PUT, DELETE, etc.
            crossDomain: true,
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            console.log(response.headers)
            alert()

            return response.status
        } else {
            alert(response.status)
            return response.status
        }
    },
    //like////////////////////////////////////////////////
    //Get
    async likePostGet(id) {
        let response = await fetch("http://localhost:8080/like/post/" + id);

        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            let json = await response.json();
            return json
        } else {
            return "Ошибка HTTP: " + response.status
        }
    },
    async likeCommentGet(id) {
        let response = await fetch("http://localhost:8080/like/comment/" + id);

        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            let json = await response.json();
            return json
        } else {
            return "Ошибка HTTP: " + response.status
        }
    },
    //Post
    async likePost(data, id) {
        const response = await fetch("http://localhost:8080/like/post/add/" + id, {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            crossDomain: true,
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            console.log(response.headers)
            alert()

            return response.status
        } else {
            alert(response.status)
            return response.status
        }
    },
    async likeComment(data, id) {
        const response = await fetch("http://localhost:8080/like/comment/add/" + id, {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            crossDomain: true,
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            console.log(response.headers)
            alert()

            return response.status
        } else {
            alert(response.status)
            return response.status
        }
    },
    async dislikePost(data, id) {
        const response = await fetch("http://localhost:8080/dislike/post/add/" + id, {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            crossDomain: true,
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            console.log(response.headers)
            alert()

            return response.status
        } else {
            alert(response.status)
            return response.status
        }
    },
    async dislikeComment(data, id) {
        const response = await fetch("http://localhost:8080/dislike/comment/add/" + id, {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            crossDomain: true,
            headers: {
                'Content-Type': 'application/json'
                    // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify(data) // body data type must match "Content-Type" header
        });
        if (response.ok) { // если HTTP-статус в диапазоне 200-299
            // получаем тело ответа (см. про этот метод ниже)
            console.log(response.headers)
            alert()

            return response.status
        } else {
            alert(response.status)
            return response.status
        }
    },
};