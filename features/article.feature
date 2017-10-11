Feature: Ability to search articles offline

Scenario: Consume article from elastic
    Given I have an article with details
    |author  |title                                 |description                                               |url                                             |urlToImage                                                                                  |publishedAt         |
    |BBC News|Schools occupied before Catalonia vote|Police move against schools designated as polling station |http://www.bbc.co.uk/news/world-europe-41452174 |https://ichef.bbci.co.uk/news/1024/cpsprodpb/7BA0/production/_98084613_mediaitem98084612.jpg|2017-09-30T17:46:26Z|
    Then I query elastic
    And I get status "200"



