WITH o as(select DISTINCT user_id from orders order by user_id)

select u.id, u.name from o 
LEFT join users as u ON u.id = o.user_id
GROUP by u.id 
ORDER by u.name, u.id
