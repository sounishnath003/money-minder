package core

const (
	QUERY_GET_TRANSACTIONS_BY_USER_ID = `SELECT 
	ID, Name, TransactionType, CategoryID, UserID, Amount, CreatedAt 
	FROM sounish-cloud-workstation.money_minder.t_transactions 
	WHERE UserID=@UserID AND CreatedAt BETWEEN @FromDate AND @EndDate
	ORDER BY CreatedAt DESC;`

	QUERY_GET_ACTIVE_CATRGORY = `SELECT 
	ID, Category FROM sounish-cloud-workstation.money_minder.t_categories
	WHERE IsActive = true ORDER BY ID ASC;`

	QUERY_ADD_NEW_TRANSACTION = `
	INSERT INTO sounish-cloud-workstation.money_minder.t_transactions
    (ID,Name,TransactionType,UserID,CategoryID,Amount,CreatedAt)
    VALUES (@ID, @Name, @TransactionType, @UserID, @CategoryID, @Amount, @CreatedAt);`
)
