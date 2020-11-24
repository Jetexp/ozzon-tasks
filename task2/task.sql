SELECT id, name
FROM goods
LEFT JOIN tags_goods ON tags_goods.goods_id=goods.id
WHERE tags_goods.goods_id IS NOT NULL
AND (SELECT COUNT(id) FROM tags) = (
        SELECT COUNT(tag_id)
            FROM tags_goods WHERE id = tags_goods.goods_id
        )
GROUP BY id, name;