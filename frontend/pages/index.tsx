// frontend/pages/index.tsx
import { useState } from "react";
import { fetchJSON } from "../lib/api";
import { postJSON } from "../lib/api";
import { putJSON } from "../lib/api";
import { deleteJSON } from "../lib/api";

type Paper = {
  id: number;
  title: string;
  authors: string;
  conference: string;
  year: number;
  url: string;
  citation_count: number;
  abstract: string;
  bibtex: string;
  pdf_url: string;
  created_at: string;
  updated_at: string;
};

export default function HomePage() {
  const [title, setTitle] = useState("");
  const [year, setYear] = useState("");
  const [results, setResults] = useState<Paper[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleSearch = async () => {
    const params = new URLSearchParams();
    if (title.trim()) params.append("title", title.trim());
    if (year.trim()) params.append("year", year.trim());

    if (!params.toString()) return;

    setLoading(true);
    setError("");
    try {
      const res = await fetchJSON<{ papers: Paper[] }>(`/papers?${params.toString()}`);
      setResults(res.papers);
    } catch (err: any) {
      setError(err.message || "検索中にエラーが発生しました");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <h1 className="text-2xl font-bold mb-6 text-black">論文検索</h1>

      <div className="flex flex-col sm:flex-row gap-4 mb-8">
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="タイトルで検索"
          className="border px-4 py-2 rounded shadow-sm flex-1 text-gray-700"
        />
        <input
          type="number"
          value={year}
          onChange={(e) => setYear(e.target.value)}
          placeholder="発表年"
          className="border px-4 py-2 rounded shadow-sm w-40 text-gray-700"
        />
        <button
          onClick={handleSearch}
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        >
          検索
        </button>
      </div>

      {loading && <p className="text-gray-600 mb-4">検索中...</p>}
      {error && <p className="text-red-500 mb-4">{error}</p>}

      <div className="space-y-4">
        {results.map((paper) => (
          <div key={paper.id} className="p-4 border rounded bg-white shadow-sm">
            <div className="font-bold text-lg text-gray-800">{paper.title}</div>
            <div className="text-sm text-gray-600">{paper.authors}</div>
            <div className="text-sm text-gray-500">
              {paper.conference} / {paper.year}
            </div>
            <a
              href={paper.url}
              target="_blank"
              rel="noopener noreferrer"
              className="text-blue-600 text-sm underline"
            >
              論文リンク
            </a>
          </div>
        ))}
      </div>
    </div>
  );
}
