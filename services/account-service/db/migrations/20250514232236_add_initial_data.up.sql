INSERT INTO permission (id, code, parent_id) VALUES (1, 'SAR', null),
                                                    (2, 'LR', 1),
                                                    (3, 'CR', 1),
                                                    (4, 'UR', 1),
                                                    (5, 'DR', 1),
                                                    (6, 'SAU', null),
                                                    (7, 'LU', 6),
                                                    (8, 'CU', 6),
                                                    (9, 'UU', 6),
                                                    (10, 'DU', 6),
                                                    (11, 'LSA', null);

INSERT INTO module (id, code) VALUES (1, 'AC');