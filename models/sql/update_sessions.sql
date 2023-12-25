UPDATE sessions 
SET token_hash = 'ABC123'
WHERE user_id = 1
RETURNING id 
