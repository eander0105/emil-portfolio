import fs from 'fs';
import path from 'path';

// TODO: make this function save file in db (look in `api/image/[id]/+server.ts`)
export async function POST({ request }: { request: Request }): Promise<Response> {
	const data = await request.formData();
	const file = data.get('file');

	if (!(file instanceof File)) {
		return new Response(JSON.stringify({ error: 'No file uploaded or invalid file type' }), { status: 400 });
	}

	const uploadDir = 'static/uploads';
	if (!fs.existsSync(uploadDir)) fs.mkdirSync(uploadDir, { recursive: true });

	const filePath = path.join(uploadDir, file.name);
	try {
		const buffer = Buffer.from(await file.arrayBuffer());
		fs.writeFileSync(filePath, buffer);
	} catch (err) {
		return new Response(JSON.stringify({ error: 'Failed to save file' }), { status: 500 });
	}

	return new Response(JSON.stringify({ success: true, path: `/uploads/${file.name}` }), { status: 200 });
}
