package core

const (
	GET_TRANSACTIONS_BY_USER_ID = `SELECT ID, Name, TransactionType, CategoryID, UserID, CreatedAt FROM sounish-cloud-workstation.money_minder.t_transactions WHERE UserID=@UserID`
)
