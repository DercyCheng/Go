-- 题1：用户在自己好友中的排名（同分并列）
WITH valid_friends AS (
    SELECT user_id, friend_id
    FROM t_friend
    WHERE user_id != friend_id  -- 排除自己加自己的无效数据
    )
SELECT
    u.user_id,
    f.friend_date AS date,
  DENSE_RANK() OVER (
    PARTITION BY u.user_id, u.date
    ORDER BY u.step DESC
  ) AS friend_rank
FROM t_user_step u
    JOIN valid_friends f
ON u.user_id = f.user_id  -- 关联有效好友
    JOIN t_user_step friend_step
    ON f.friend_id = friend_step.user_id
    AND u.date = friend_step.date  -- 关联好友步数

-- 题2：好友在用户数据中的排名（同分并列）
WITH valid_friends AS (
    SELECT user_id, friend_id
    FROM t_friend
    WHERE user_id != friend_id  -- 过滤自循环数据
    )
SELECT
    f.user_id,
    u.date,
    DENSE_RANK() OVER (
    PARTITION BY f.friend_id, u.date
    ORDER BY u.step DESC
  ) AS user_rank
FROM t_user_step u
         JOIN valid_friends f
              ON u.user_id = f.friend_id  -- 反向关联用户数据
         JOIN t_user_step self
              ON f.user_id = self.user_id
                  AND u.date = self.date  -- 关联用户自身步数

-- 题 3：用户占据好友封面的次数（每日统计）
WITH valid_friends AS (
  SELECT
    user_id,
    friend_id,
    date
  FROM t_friend
  WHERE user_id != friend_id  -- 过滤自关联
),

-- 获取所有将用户作为好友的宿主用户（host_user）
host_relationships AS (
  SELECT
    user_id AS host_user,
    friend_id AS user_id,
    date
  FROM valid_friends
),

-- 获取宿主用户当天的所有好友及其步数
host_friend_steps AS (
  SELECT
    h.host_user,
    h.user_id AS cover_user,
    h.date,
    f.friend_id,
    s.step_cnt
  FROM host_relationships h
  JOIN valid_friends f       -- 获取宿主用户的好友列表
    ON h.host_user = f.user_id
    AND h.date = f.date
  JOIN t_step s              -- 关联好友步数
    ON f.friend_id = s.user_id
    AND f.date = s.date
),

-- 计算宿主用户当日好友排名
friend_rankings AS (
  SELECT
    host_user,
    date,
    friend_id,
    DENSE_RANK() OVER (
      PARTITION BY host_user, date
      ORDER BY step_cnt DESC
    ) AS rank
  FROM host_friend_steps
)

-- 统计用户占据封面的次数
SELECT
    friend_id AS user_id,
    date,
    COUNT(*) AS cover_count
FROM friend_rankings
WHERE rank = 1                -- 只取封面（第一名）
GROUP BY friend_id, date