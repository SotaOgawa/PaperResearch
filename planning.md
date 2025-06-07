# 開発計画書 - 論文管理アプリ（Go + Next.js）

## 開発目的

- Web エンジニア職を志望するにあたり、バックエンド・フロントエンド・DB を一貫して構築できる力を示す
- 研究室で実際に使える論文管理アプリとしての実用性も追求

## 開発メンバー

- 1 人（チーム想定で役割分担を意識）

| 役割                 | 担当内容                       |
| -------------------- | ------------------------------ |
| プロジェクトオーナー | ユーザー視点で機能定義         |
| スクラムマスター     | スケジュール・進行管理         |
| 開発者               | 設計・実装全般（Go + Next.js） |

## スクラム的進行

- 各スプリントで「設計 → 実装 → 動作確認 → 振り返り」を実施

| Sprint   | 内容                                |
| -------- | ----------------------------------- |
| Sprint 1 | 要件定義・API 設計・DB 設計         |
| Sprint 2 | Go API（CRUD）実装                  |
| Sprint 3 | Next.js フロント（一覧・登録）      |
| Sprint 4 | BibTeX/PDF 自動取得実装・統合       |
| Sprint 5 | デプロイ・README 整備・提出用まとめ |

## プロダクトに関する情報

### プロダクトビジョン

研究室で用いている論文をまとめた db をより参照しやすくし、同時にエンジニアとしてのスキルを示すためのポートフォリオ

### プロダクトバックログ

 | 項目名                                                                                   | 優先度 | 完了 |
  | ---------------------------------------------------------------------------------------- | ------ | ---- |
  | 論文一覧表示できるようにしたい                                                           | 5      | Done |
  | 検索できるようにしたい                                                                   | 5      | Title, Yearのみ |
  | 簡単に登録できるように、一括で登録できるようにしたい                                     | 5      |Done |
  | 著者データは正規化できるようにしたい                                                     | 3      | Rejected |
  | フロントから簡単に論文が検索できるようにしたい                                           | 3      | Done |
  | bibtex や PDF も取得できるようにしたい                                                   | 3      |Done |
  | abstract や title とキーワード、方向性のコサイン類似度などでマッチングできるようにしたい | 4      | |
  | DBの未入力項目に「未入力」印を表示したい、これを検知してSemantic Scholar側で対応するようにしたい               | 3      | Done |
| フロントで論文項目をクリックすると詳細が展開されるようにしたい                         | 4      | Done |
| 開発者用画面を用意し、Semantic Scholar API のレスポンス確認や手動クロール等を行えるようにしたい | 3〜4    | |

## 使用技術

- Go
- SQLite + GORM
- Next.js + TailwindCSS

## DONE リスト（随時更新）

- [x] API & DB 設計
- [x] Go: CRUD API 実装
- [x] Next.js: UI 実装
- [x] 実際の論文情報の更新
- [ ] デプロイ完了

## 各種設計

### DB構造

#### papersテーブル

| カラム名       | 型       | 説明                  |
| -------------- | -------- | --------------------- |
| id             | INTEGER  | 論文ID                |
| title          | TEXT     | 論文タイトル          |
| conference     | TEXT     | 学会名                |
| year           | INTEGER  | 発表年                |
| authors        | TEXT     | 著者(,で結合)         |
| abstract       | TEXT     | 概要                  |
| url            | TEXT     | リンク先 URL          |
| citation_count | INTEGER  | 引用数                |
| bibtex         | TEXT     | BibTeX 文字列         |
| pdf_url        | TEXT     | PDFリンク or 保存パス |
| updated_at     | DATETIME | 更新日時              |
| created_at     | DATETIME | 登録日時              |

### API構造

  | メソッド | エンドポイント    | 内容                         |
  | -------- | ----------------- | ---------------------------- |
  | GET      | `/api/papers`     | 論文一覧の取得               |
  | POST     | `/api/papers`     | 論文の新規登録|
  | PUT      | `/api/papers/:id` | 論文の更新                   |
  | DELETE   | `/api/papers/:id` | 論文の削除                   |

## 開発履歴

### Sprint 1

- **プロダクトビジョン**: 研究室で用いている論文をまとめた db をより参照しやすくし、同時にエンジニアとしてのスキルを示すためのポートフォリオ
- **プロダクトバックログアイテム**

  | 項目名                                                                                   | 優先度 | 完了 |
  | ---------------------------------------------------------------------------------------- | ------ | ---- |
  | 論文一覧表示できるようにしたい                                                           | 5      | Done |
  | 検索できるようにしたい                                                                   | 5      | Title, Yearのみ |
  | 簡単に登録できるように、一括で登録できるようにしたい                                     | 5      |Done |
  | 著者データは正規化できるようにしたい                                                     | 3      | Rejected |
  | フロントから簡単に論文が検索できるようにしたい                                           | 3      | |
  | bibtex や PDF も取得できるようにしたい                                                   | 3      |Done |
  | abstract や title とキーワード、方向性のコサイン類似度などでマッチングできるようにしたい | 4      | |

- **スプリント計画**

| Sprint   | 内容                                |
| -------- | ----------------------------------- |
| Sprint 1 | 要件定義・API 設計・DB 設計         |
| Sprint 2 | Go API（CRUD）実装                  |
| Sprint 3 | Next.js フロント（一覧・登録）      |
| Sprint 4 | BibTeX/PDF 自動取得実装・統合       |
| Sprint 5 | デプロイ・README 整備・提出用まとめ |

- **完成要件**

  - 設計完了を「API 仕様書」「DB 設計」「Git での管理体制」とする
  - 実装完了を「コード実装完了」「カバレッジ 90%以上のテストがあり、これが動く」「README が更新されている」とする

#### api設計

  | メソッド | エンドポイント    | 内容                         |
  | -------- | ----------------- | ---------------------------- |
  | GET      | `/api/papers`     | 論文一覧の取得               |
  | GET      | `/api/papers/:id` | 論文詳細の取得               |
  | POST     | `/api/papers`     | 論文の新規登録（自動取得含） |
  | PUT      | `/api/papers/:id` | 論文の更新                   |
  | DELETE   | `/api/papers/:id` | 論文の削除                   |

#### db 設計

##### papers テーブル

| カラム名       | 型       | 説明                  |
| -------------- | -------- | --------------------- |
| id             | INTEGER  | 論文ID                |
| title          | TEXT     | 論文タイトル          |
| conference     | TEXT     | 学会名                |
| year           | INTEGER  | 発表年                |
| authors        | TEXT     | 著者(,で結合)         |
| abstract       | TEXT     | 概要                  |
| url            | TEXT     | リンク先 URL          |
| citation_count | INTEGER  | 引用数                |
| bibtex         | TEXT     | BibTeX 文字列         |
| pdf_url        | TEXT     | PDFリンク or 保存パス |
| updated_at     | DATETIME | 更新日時              |
| created_at     | DATETIME | 登録日時              |

##### authors テーブル

| カラム名  | 型       | 説明                 |
| ----- | ------- | ------------------ |
| id    | INTEGER | 著者ID               |
| name  | TEXT    | 著者名                |
| orcid | TEXT    | ORCID識別子（optional） |

##### paper_authors テーブル

| カラム名          | 型       | 説明   |
| ------------- | ------- | ---- |
| paper\_id     | INTEGER | 論文ID |
| author\_id    | INTEGER | 著者ID |
| author\_order | INTEGER | 著者順  |

#### Sprint完了

- Sprint Review
  - API仕様書完了, DB設計完了
  - バックログある程度完了
  - Gitの初期環境を構築済み
- Sprint Retrospective
  - 次のSprintへと順調に進められる
  - Go、Next.jsに理解を深めながら実装していきたい

### Sprint 2

#### DoD

- APIの実装、機能するものを作成すること
- SQLiteを用いたローカルDBでデータが永続化されている
- curlでの動作確認
- README.mdに使い方が記述されている
- テストケースが書かれている

#### API設計

- ID検索とエンドポイントとして分ける必要なし、GETメソッドに一本化する

  | メソッド | エンドポイント    | 内容                         |
  | -------- | ----------------- | ---------------------------- |
  | GET      | `/api/papers`     | 論文一覧の取得               |
  | POST     | `/api/papers`     | 論文の新規登録（自動取得含） |
  | PUT      | `/api/papers/:id` | 論文の更新                   |
  | DELETE   | `/api/papers/:id` | 論文の削除                   |
  
- 実装完了

#### SQLiteDBの実装, curlでの動作確認

- テスト済み

#### テストケース

- すべてのAPI（GET, POST, PUT, DELETE）に対して単体テストを作成
- Go + httptest + sqlite (in-memory) 、およびDBの注入により依存の少ない高速テストを実現
- テスト実行:

```bash
go test ./... -cover
```

#### Sprint完了

- Sprint Review
  - API完成、テストケース完成
  - カバレッジは60%程度なので今後余裕を見つけて拡充
- Sprint Retrospective
  - 今のところかなり順調
  - CI/CDの導入を考えたい

### Sprint 3

#### CI/CD導入

- 一旦見送り

#### pingの導入

- 面白そうなのでping/pongを実装した。

#### DBの構図を変更

- 現在用いる予定であるOpenReviewのAPIでは著者の固有IDが取得可能なので、Authorテーブルにそれを用いてIdenticalな情報を付与可能
- しかし他の手法で取得する論文(ASIACCS)ではこの方法が使えない、Authorが被った際の対処ができない
- **著者を一人ずつ管理するのをやめ、Authorsとして一つの項目に押し込む**
- 欠点: 著者での検索が少々厄介になるが仕方がない、完成を目指す

##### paper テーブル

| カラム名       | 型       | 説明                  |
| -------------- | -------- | --------------------- |
| id             | INTEGER  | 論文ID                |
| title          | TEXT     | 論文タイトル          |
| conference     | TEXT     | 学会名                |
| year           | INTEGER  | 発表年                |
| authors        | TEXT     | 著者(,で結合)         |
| abstract       | TEXT     | 概要                  |
| url            | TEXT     | リンク先 URL          |
| citation_count | INTEGER  | 引用数                |
| bibtex         | TEXT     | BibTeX 文字列         |
| pdf_url        | TEXT     | PDFリンク or 保存パス |
| updated_at     | DATETIME | 更新日時              |
| created_at     | DATETIME | 登録日時              |

#### crawlerの実装

- OpenReviewからダウンロードされる形式、DBに保存する構造などを一つずつ定義した
- まず、OpenReviewから取得できる情報のうち、Title, Authors, Year, Conferenceのみを登録するRawPaperテーブルを保存。
- 次に外部APIを用いてabstractなどを逐次保存していく。
- ICMLを試しにRawPaper形式で登録できるようにして、動作を確認

#### Next.jsを用いたフロントの登録

- GET /api/papers?title=...&year=... に対応した検索画面を構築
- 論文タイトル・発表年によるフィルタ入力フォームを作成
- 検索結果をリスト形式で表示（タイトル・著者・学会名・年・リンク）
- TailwindCSS によるスタイリング

#### テストの実装

- Jest + testing libraryによるユニットテストの作成
- モックによってNext.js側のみのユニットテストを作成

#### Sprint完了

- Sprint Review
  - Next.js による一覧・検索画面が完成
  - 最低限動くシステムの完成
  - TailwindCSSにより見た目も整備
  - Jestテストが通る状態で信頼性を担保
- Sprint Retrospective
  - フロントとバックの連携が想定通りに動作した
  - テスト・スタイリング・状態管理など一通り実装できたため、次Sprintはより具体的なシステム設計に着手する。

### Sprint 4

#### WoSのAPIを用いた追加データの拡張

- **WoSではabstractを取得できない**ことが判明(学校の提供するAPIのグレードが低い)
- 代替APIとして**Semantic Scholar**を利用
- SSのAPIを用いたabstractなどの取得まで完了

#### バックログアイテムの拡充

| 項目名 | 優先度 |
| ---- | --- |
| DBの未入力項目に「未入力」印を表示したい、これを検知してSemantic Scholar側で対応するようにしたい                                                   | 3      |
| フロントで論文項目をクリックすると詳細が展開されるようにしたい                         | 4      |
| 開発者用画面を用意し、Semantic Scholar API のレスポンス確認や手動クロール等を行えるようにしたい | 3〜4    |
| 検索結果をcsvで出力する機能 | 4 |

#### フロントで論文項目をクリックすると詳細が展開されるようにしたい

- componentsとしてPaperModalを実装
- Modal外のクリック、およびバツボタンのクリックによって閉じられる設計にした
- 背景の透明化をbg-black/50で指定した

#### Sprint完了

- Sprint Review
  - Semantic Scholar APIによる補完機能を導入。WoSではabstract取得不可のため、SSに切替
  - 未入力項目の判定とUI表示（"未入力"印）を設計
  - 論文詳細モーダル（PaperModal）を実装
    - 論文一覧項目クリック → ポップアップ表示
    - 背景透過（bg-black/50）でUX向上
    - ×ボタン、外クリックで閉じる動作に対応
  - 最低限の動作確認まで完了
- Sprint Retrospective
  - Tailwindのopacityバグや背景描画問題を適切に切り分けて解決できた
  - UI/UXの細部（モーダル、クリック判定、透過）に気を配り、完成度が向上した
  - 外部APIの切り替え判断が早く、進行に支障が出なかった点が良かった
  - ただし、補完処理の自動化（CLI/管理画面）、開発者画面、CSV出力などは時間不足で未完。

### Sprint 5

#### デプロイとポートフォリオ整備

- Go バックエンドを Railway にてデプロイ（PostgreSQL に対応）
- Next.js フロントエンドを Vercel にてデプロイ
- .env を使った環境変数管理と CORS 設定によるフロント連携
- Semantic Scholar API の導入により論文情報の補完を実現
- ON CONFLICT による重複挿入の抑制、および GORM マイグレーション設計
- UNIQUE 制約を手動で追加し、衝突検知と更新処理を実装
- GitHub に公開可能な状態で README.md を再構成（目的・使用技術・アピールポイントを明示）
- CLI や GUI（Railway）経由で PostgreSQL を操作可能な状態を整備
- papers.db から PostgreSQL への完全移行を完了

#### DoD

- 本番環境で動作する Go + Next.js アプリケーションがデプロイされている
- Semantic Scholar によるデータ拡張が動作する
- README.md がポートフォリオ提出に耐える水準まで整備されている
- 本番用 GitHub リポジトリが整理されている

#### Sprint完了

- Sprint Review
  - バックエンド・フロントエンドともにクラウド上で動作する形で統合が完了
  - データベースは PostgreSQL に統一、Semantic Scholar による情報拡張も動作確認済
  - README.md を就活向けに整備し、自己紹介・目的・デプロイURL・技術構成を一貫して記載
- Sprint Retrospective
  - Railway, Vercel などの実運用に近い構成を経験できたのは大きな成果
  - DB移行やCORS、UNIQUE制約、APIエラー処理などバックエンド特有の実装課題に対応できた
  - 今後はLLM連携・データ可視化・ユーザー管理などを検討予定