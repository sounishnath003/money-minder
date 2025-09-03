package core

const (
	QUERY_GET_TRANSACTIONS_BY_USER_ID = `SELECT 
	ID, Name, TransactionType, PaymentMethod, CategoryID, UserID, Amount, CreatedAt 
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
		AND b.Category NOT IN ('Salary')
  	GROUP BY b.Category
	ORDER BY TotalAmount DESC, Category ASC
	LIMIT 6;`

	QUERY_GET_ACTIVE_CATRGORY = `SELECT 
	ID, Category FROM sounish-cloud-workstation.money_minder.t_categories
	WHERE IsActive = true ORDER BY ID ASC;`

	QUERY_ADD_NEW_TRANSACTION = `
	INSERT INTO sounish-cloud-workstation.money_minder.t_transactions
    (ID,Name,TransactionType,UserID,CategoryID,PaymentMethod,Amount,CreatedAt)
    VALUES (@ID, @Name, @TransactionType, @UserID, @CategoryID, @PaymentMethod, @Amount, @CreatedAt);`

	ANALYTICS_TOTAL_DAILY_SPEND_BY_DATE_RANGE = `SELECT
	UNIX_MILLIS(TIMESTAMP(DATE(CreatedAt))) AS UnixMiliseconds, SUM(Amount) AS Amount
	FROM sounish-cloud-workstation.money_minder.t_transactions
	WHERE UserID=@UserID AND TransactionType = 'Expense' AND CreatedAt BETWEEN @FromDate AND @EndDate
	GROUP BY UnixMiliseconds ORDER BY UnixMiliseconds;
	`

	ANALYTICS_SPENDS_ON_CATEGORIES_MONTH_ON_MONTH = `WITH a AS (
	SELECT CreatedAt, CategoryID, Amount FROM sounish-cloud-workstation.money_minder.t_transactions
	WHERE UserID=@UserID AND CAST(CreatedAt AS DATE) BETWEEN DATE_SUB(CURRENT_DATE(), INTERVAL 4 MONTH) AND CURRENT_DATE())
	SELECT FORMAT_DATE("%m%Y", a.CreatedAt) AS Seq, FORMAT_DATE("%b-%Y", a.CreatedAt) AS Month, 
	b.Category AS Category, SUM(a.Amount) AS TotalSpendAmount
	FROM a JOIN sounish-cloud-workstation.money_minder.t_categories AS b 
	ON a.CategoryID = b.ID WHERE b.IsActive = true 
	GROUP BY Seq, Month, Category 
	ORDER BY Seq ASC;`

	ANALYTICS_SPENDS_ON_CATEGORIES_BY_ALL_YEARMONTH_AGGREGATED = `WITH monthlyExpenses AS (SELECT EXTRACT(YEAR FROM CreatedAt) AS Year, EXTRACT(MONTH FROM CreatedAt) AS Month, FORMAT_DATE("%b-%Y", CreatedAt) AS MonthYearStr, CategoryID, PaymentMethod, SUM(Amount) AS TotalAmount FROM sounish-cloud-workstation.money_minder.t_transactions WHERE UserID=@UserID AND CAST(CreatedAt AS DATE) BETWEEN DATE_SUB(CURRENT_DATE(), INTERVAL 12 MONTH) AND CURRENT_DATE() GROUP BY Year, Month, MonthYearStr, CategoryID, PaymentMethod ) SELECT m.Year, m.Month, m.MonthYearStr, b.Category, m.PaymentMethod, m.TotalAmount FROM monthlyExpenses AS m JOIN sounish-cloud-workstation.money_minder.t_categories AS b ON m.CategoryID = b.ID WHERE b.IsActive = TRUE ORDER BY m.Year DESC, m.Month DESC, m.TotalAmount DESC;`

	ANALYTICS_MONTHLY_SAVINGS_BREAKDOWN = `WITH base AS (
    SELECT 
      EXTRACT(YEAR FROM a.CreatedAt) AS Year,
      EXTRACT(MONTH FROM a.CreatedAt) AS Month,
      FORMAT_DATE("%b-%Y", a.CreatedAt) AS MonthYearStr,
      b.Category,
      a.Amount
    FROM sounish-cloud-workstation.money_minder.t_transactions AS a
    JOIN sounish-cloud-workstation.money_minder.t_categories AS b 
      ON a.CategoryID = b.ID
    WHERE a.UserID = @UserID 
      AND b.IsActive = TRUE
      AND CAST(a.CreatedAt AS DATE) BETWEEN DATE_SUB(CURRENT_DATE(), INTERVAL 12 MONTH) AND CURRENT_DATE()
  ), monthly AS (
    SELECT 
      Year,
      Month,
      MonthYearStr,
      -- Only Salary counts as income
      SUM(CASE WHEN Category = 'Salary' THEN Amount ELSE 0 END) AS TotalIncome,
      -- Mutual Funds treated as savings
      SUM(CASE WHEN Category = 'Mutual Funds' THEN Amount ELSE 0 END) AS Investments,
      -- All other categories are expenses
      SUM(CASE WHEN Category NOT IN ('Salary', 'Mutual Funds') THEN Amount ELSE 0 END) AS RawExpense
    FROM base
    GROUP BY Year, Month, MonthYearStr
  ), computed AS (
    SELECT
      Year,
      Month,
      MonthYearStr,
      TotalIncome,
      RawExpense,
      Investments,
      RawExpense AS TotalExpense,
      -- Net savings = Salary income - expenses + mutual funds investments
      (TotalIncome - RawExpense + Investments) AS NetSavings
    FROM monthly
  )
  SELECT 
    Year,
    Month,
    MonthYearStr,
    TotalIncome,
    TotalExpense,
    -- expose breakdown for UI toggling
    RawExpense,
    Investments,
    NetSavings,
    CASE 
      WHEN LAG(NetSavings) OVER (ORDER BY Year, Month) IS NULL THEN 0
      ELSE NetSavings - LAG(NetSavings) OVER (ORDER BY Year, Month)
    END AS SavingsChange,
    CASE 
      WHEN LAG(NetSavings) OVER (ORDER BY Year, Month) IS NULL OR LAG(NetSavings) OVER (ORDER BY Year, Month) = 0 THEN 0
      ELSE ROUND(((NetSavings - LAG(NetSavings) OVER (ORDER BY Year, Month)) / ABS(LAG(NetSavings) OVER (ORDER BY Year, Month))) * 100, 2)
    END AS SavingsChangePercent
  FROM computed
  ORDER BY Year DESC, Month DESC;`
)
