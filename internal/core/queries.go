package core

const (
	QUERY_GET_TRANSACTIONS_BY_USER_ID = `SELECT 
	ID, Name, TransactionType, CategoryID, UserID, Amount, CreatedAt 
	FROM sounish-cloud-workstation.money_minder.t_transactions 
	WHERE UserID=@UserID AND CreatedAt BETWEEN @FromDate AND @EndDate
	ORDER BY CreatedAt DESC;`

	QUERY_GET_TRANSACTION_SPEND_BY_CATEGORY = `SELECT 
	b.Category AS Category, SUM(a.Amount) AS TotalAmount
	FROM sounish-cloud-workstation.money_minder.t_transactions AS a
  	JOIN sounish-cloud-workstation.money_minder.t_categories AS b
  	ON a.categoryID = b.ID
	WHERE 
		UserID=@UserID 
		AND CreatedAt BETWEEN @FromDate AND @EndDate 
		AND b.IsActive = true
  	GROUP BY b.Category
	ORDER BY TotalAmount DESC, Category ASC
	LIMIT 6;`

	QUERY_GET_ACTIVE_CATRGORY = `SELECT 
	ID, Category FROM sounish-cloud-workstation.money_minder.t_categories
	WHERE IsActive = true ORDER BY ID ASC;`

	QUERY_ADD_NEW_TRANSACTION = `
	INSERT INTO sounish-cloud-workstation.money_minder.t_transactions
    (ID,Name,TransactionType,UserID,CategoryID,Amount,CreatedAt)
    VALUES (@ID, @Name, @TransactionType, @UserID, @CategoryID, @Amount, @CreatedAt);`
)
