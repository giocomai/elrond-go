package txcache

const estimatedSizeOfBoundedTxFields = uint64(128)

// estimateTxSize returns an approximation
func estimateTxSize(tx *WrappedTransaction) uint64 {
	sizeOfData := uint64(len(tx.Tx.GetData()))
	return estimatedSizeOfBoundedTxFields + sizeOfData
}

// estimateTxGas returns an approximation for the necessary computation units (gas units)
func estimateTxGas(tx *WrappedTransaction) uint64 {
	gasLimit := tx.Tx.GetGasLimit()
	return gasLimit
}

// estimateTxFee returns an approximation for the cost of a transaction, in nano ERD
// TODO: switch to integer operations (as opposed to float operations).
// TODO: do not assume the order of magnitude of minGasPrice.
func estimateTxFee(tx *WrappedTransaction) uint64 {
	// In order to obtain the result as nano ERD (not as "atomic" 10^-18 ERD), we have to divide by 10^9
	// In order to have better precision, we divide the factors by 10^6, and 10^3 respectively
	gasLimit := float32(tx.Tx.GetGasLimit()) / 1000000
	gasPrice := float32(tx.Tx.GetGasPrice()) / 1000
	feeInNanoERD := gasLimit * gasPrice
	return uint64(feeInNanoERD)
}