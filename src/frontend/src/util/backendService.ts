const BACKEND_PORT = 8081;
const BACKEND_URL = `http://127.0.0.1:${BACKEND_PORT}`;

export async function fetchFromBackend(request: Request, endpoint: string, data?: object): Promise<Response> {
	const cookie = request.headers.get('cookie');

	const headers: HeadersInit = {
		'Accept': 'application/json',
		'Content-Type': 'application/json'
	};

	if(cookie) {
		headers['Cookie'] = cookie;
	}

    const response = await fetch(`${BACKEND_URL}/${endpoint}`, {
        method: 'POST',
		credentials: 'include',
        headers: headers,
        body: JSON.stringify(data)
    });

    if (response.ok) {
        return response;
    } else {
        console.error('Error:', response.statusText);
    }

    return Promise.reject("The request to the backend failed to resolve.");
}