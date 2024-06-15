//const url = "http://192.168.122.40:8000";
//const url = "http://localhost:8000";
export function getRam() {
    return fetch(`/insertRam`, {
        headers: {'Content-Type': 'application/json'},
        method: 'GET',
    });
}


export function getCPU() {
    return fetch(`/InsertCPU`, {
        headers: {'Content-Type': 'application/json'},
        method: 'GET',
    });
}



export function createProcess() {
    return fetch(`/insertProcess`, {
        method: 'GET',
        headers: {'Content-Type': 'application/json'},
    })
}

export function deleteProcess(pid) {
    return fetch(`/killProcess?pid=${pid}`, {
        method: 'GET',
        headers: {'Content-Type': 'application/json'},
    })
}
