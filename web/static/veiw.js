export default {
    render(tempName, model) {
        const tempsElement = document.getElementById(tempName);
        const tempSours = tempsElement.innerHTML;
        const rendrFn = Handlebars.compile(tempSours)
        return rendrFn(model)
    }
}