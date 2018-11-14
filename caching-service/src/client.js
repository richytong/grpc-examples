const grpc = require('grpc')
const messages = require('../proto/cache_pb')
const services = require('../proto/cache_grpc_pb')

const stringToBytes = (bufferString) => {
	const arr = new Uint8Array(bufferString.length);
	const stringCharCodes = bufferString.split('').map((_, idx) => bufferString.charCodeAt(idx))
	arr.set(stringCharCodes)
	return arr;
}

const handleResp = (err, resp) => {
	if (err) {
		throw err
	} else {
		console.log(`success: ${resp}`)
	}
}

const main = async () => {
	const client = new services.CacheClient(
		'localhost:5051',
		grpc.credentials.createInsecure()
	)
	let req = new messages.SetReq()
	req.setKey('gopher')
	req.setVal(stringToBytes('con'))
	client.set(req, handleResp)
	req = new messages.GetReq()
	req.setKey('gopher')
	client.get(req, handleResp)
}

main().catch(err => console.log(err))
