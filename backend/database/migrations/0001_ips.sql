--  simple satire
CREATE TABLE death_note (
	id				SERIAL PRIMARY KEY
,	ip	  			TEXT NOT NULL
,	user_agent		TEXT
,	created_at 		DATE DEFAULT timezone('BRT', CURRENT_TIMESTAMP)
);