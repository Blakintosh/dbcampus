const BACKEND_PORT = 8081;
const BACKEND_URL = `http://127.0.0.1:${BACKEND_PORT}`;

export async function fetchFromBackend(endpoint: string, data: object): Promise<Response> {
    console.log(JSON.stringify(data));
    const response = await fetch(`${BACKEND_URL}/${endpoint}`, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });

    if (response.ok) {
        return response;
    } else {
        console.error('Error:', response.statusText);
    }

    return Promise.reject("The request to the backend failed to resolve.");
}