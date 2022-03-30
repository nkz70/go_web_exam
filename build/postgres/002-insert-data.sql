INSERT INTO users (
    id,
    last_name,
    first_name,
    age,
    gender,
    score,
    time,
    status,
    test_datetime,
    created_at,
    updated_at
) VALUES (
    1,
    'ほげ山',
    'ふが太',
    18,
    1,
    50,
    3600,
    1,
    current_timestamp,
    current_timestamp,
    current_timestamp
);

INSERT INTO questions (
    id,
    content,
    image,
    number,
    short_name,
    status,
    created_at,
    updated_at
) VALUES (
    1,
    '世界一大きな川はどこでしょう？',
    'neko.jpg',
    1,
    'big river',
    1,
    current_timestamp,
    current_timestamp
);

INSERT INTO question_answers (
    id,
    form_type,
    question_id,
    answer,
    label_position,
    label,
    created_at,
    updated_at
) VALUES (
    1,
    '1',
    1,
    'ナイル川',
    1,
    '選択肢１',
    current_timestamp,
    current_timestamp
);
