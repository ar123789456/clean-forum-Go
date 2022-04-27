import Controller from './controller.js'

function getRouterInfo() {
    const hash = location.hash ? location.hash.slice(1) : '';
    const [name, id] = hash.split('/')

    return { name, params: { id } }
}

function handleHash() {
    const { name, params } = getRouterInfo();

    if (name) {
        const routerNAme = name + 'Route';
        Controller[routerNAme](params);
    }
}

export default {
    init() {
        addEventListener('hashchange', handleHash)
        addEventListener('submit', handleHash)

        handleHash();
    }
}