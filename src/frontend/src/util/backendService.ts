const BACKEND_PORT = 8081;
const BACKEND_URL = `http://localhost:${BACKEND_PORT}`;

export async function fetchFromBackend(endpoint: string, data: object): Promise<Response> {
    try {
        const response = await fetch(`${BACKEND_URL}/${endpoint}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (response.ok) {
            return response;
        } else {
            console.error('Error:', response.statusText);
        }
    } catch(error) {
        console.error('Error:', error);
    }

    return Promise.reject("The request to the backend failed to resolve.");
}