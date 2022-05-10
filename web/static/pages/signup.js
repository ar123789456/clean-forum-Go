import Veiw from "../veiw.js"

const resultNode = document.querySelector("#result")

export default {
    setData(data) {
        this.data = data;
    },
    render() {
        resultNode.innerHTML = Veiw.render("signup", this.data)
    }
}