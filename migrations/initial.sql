create table personas(
    id serial primary key,
    name varchar,
    story varchar
);

create table utils(
    key varchar primary key,
    value varchar
);

INSERT INTO personas(name, story)
VALUES ('Napoleon Bonaparte', 'Born 15 August 1769 – 5 May 1821, later known by his regnal name Napoleon I, was a French military officer and statesman who rose to prominence during the French Revolution and led a series of successful campaigns across Europe during the French Revolutionary and Napoleonic Wars from 1796 to 1815. He was the leader of the French Republic as First Consul from 1799 to 1804, then of the French Empire as Emperor of the French from 1804 to 1814, and briefly again in 1815.'),
       ('Alexander III of Macedon', 'Born July 356 BC – 10/11 June 323 BC, also known as Alexander the Great, was a king of the ancient Greek kingdom of Macedon. He succeeded his father Philip II to the throne in 336 BC at the age of 20 and spent most of his ruling years conducting a lengthy military campaign throughout Western Asia, Central Asia, parts of South Asia, and Egypt. By the age of 30, he had created one of the largest empires in history, stretching from Greece to northwestern India. He was undefeated in battle and is widely considered to be one of history''s greatest and most successful military commanders.');

INSERT INTO utils(key, value)
VALUES ('ping', 'pong!');