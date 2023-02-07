cd src\frontend
if not exist .\node_modules (
	npm install && npm run dev
) else (
	npm run dev
)
