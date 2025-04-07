-- 题 1：
-- 口诀：按天分组，去重计数
SELECT
    DATE(time) AS day,
    COUNT(DISTINCT user_id) AS uv
FROM video_play_action
GROUP BY day;

-- 题 2：
-- 口诀：分组求和 → 按日排名 → 取前三
WITH daily_sum AS (
    SELECT
        DATE(time) AS day,
        video_id,
        SUM(play_time) AS total_time
    FROM video_play_action
    GROUP BY day, video_id
),
     ranked AS (
         SELECT
             day, video_id, total_time,
             ROW_NUMBER() OVER (PARTITION BY day ORDER BY total_time DESC) AS rk
         FROM daily_sum
     )
SELECT day, video_id, total_time
FROM ranked
WHERE rk <= 3;

-- 题 3：
-- 口诀：拆分标签 → 展开统计 → 排名筛选
WITH split_tags AS (
    SELECT
        DATE(time) AS day,
        SPLIT(user_tag, '，') AS tags,  -- 按中文逗号分割
        play_time
    FROM video_play_action
),
     unnest_tags AS (
         SELECT day, tag, play_time
         FROM split_tags, UNNEST(tags) AS tag  -- 展开为多行
     ),
     tag_sum AS (
         SELECT
             day, tag, SUM(play_time) AS total_time
         FROM unnest_tags
         GROUP BY day, tag
     ),
     ranked_tags AS (
         SELECT
             day, tag, total_time,
             ROW_NUMBER() OVER (PARTITION BY day ORDER BY total_time DESC) AS rk
         FROM tag_sum
     )
SELECT day, tag, total_time
FROM ranked_tags
WHERE rk <= 3;

-- 题 4：
-- 口诀：按用户排序 → 计算间隔 → 筛选条件
WITH user_sequence AS (
    SELECT
        user_id, time, video_id, play_time,
        LAG(time) OVER (PARTITION BY user_id ORDER BY time) AS prev_time
    FROM video_play_action
),
     filter_data AS (
         SELECT
             user_id, time, video_id, play_time,
             TIMESTAMP_DIFF(time, prev_time, SECOND) AS interval_sec
         FROM user_sequence
         WHERE prev_time IS NOT NULL
           AND TIMESTAMP_DIFF(time, prev_time, SECOND) <= 60  -- 间隔≤60秒
     )
SELECT
    FORMAT_DATE('%Y%m%d', DATE(time)) AS date,  -- 格式化日期
    user_id, video_id, play_time
FROM filter_data;