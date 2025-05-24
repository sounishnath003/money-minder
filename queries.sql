
-- Insert User
    INSERT INTO
    `gcp_project_id.money_minder.t_users` (ID,
        Name,
        Email,
        Password,
        CreatedAt)
    VALUES
    (
        1, 'Sounish Nath', 'flock.sinasini@gmail.com', 'xxxxx', CURRENT_TIMESTAMP()
    );

-- Insert Category
    INSERT INTO `gcp_project_id.money_minder.t_categories`
    (ID, Category, IsActive)
    VALUES
    (8, 'xx', true)

-- Add Transaction
    INSERT INTO `gcp_project_id.money_minder.t_transactions`
    (ID,Name,TransactionType,UserID,CategoryID,Amount,CreatedAt)
    VALUES
    (2, 'Mumbai House', 'Expense', 1, 10, 14000, current_timestamp());


-- safe delete transactions (tx)
    DELETE FROM `gcp_project_id.money_minder.t_transactions`
    WHERE ID > 1747761819147;


-- MoM Income/Expense
    WITH a AS (SELECT
    *
    FROM
    `gcp_project_id.money_minder.t_transactions`
    WHERE
    UserID=1
    -- AND TransactionType = 'Expense'
    AND CreatedAt BETWEEN '2025-05-01'
    AND '2025-05-31')
    SELECT 
    FORMAT_DATE("%m-%Y",CreatedAt) AS Month, b.Category AS Category, SUM(a.Amount) AS TotalSpendAmount
    FROM a JOIN `gcp_project_id.money_minder.t_categories` AS b
    ON a.CategoryID = b.ID
    WHERE b.IsActive = true
    GROUP BY Month, Category;