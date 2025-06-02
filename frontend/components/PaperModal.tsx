import React from 'react';

interface Props {
  paper: any;
  onClose: () => void;
}

export default function PaperModal({ paper, onClose }: Props) {
  return (
    <div className="fixed inset-0 z-40 flex items-center justify-center bg-black/50">
      <div className="bg-white z-50 rounded-xl shadow-lg p-6 w-full max-w-2xl">
        <button onClick={onClose} className="absolute top-2 right-2 text-gray-500 hover:text-black text-xl font-bold">
          ×
        </button>
        <h2 className="text-xl font-bold mb-4">{paper.title}</h2>
        <div className="space-y-2 text-sm">
          <p><strong>著者:</strong> {paper.authors || '未入力'}</p>
          <p><strong>学会:</strong> {paper.conference || '未入力'}</p>
          <p><strong>年:</strong> {paper.year}</p>
          <p><strong>Abstract:</strong> {paper.abstract || '未入力'}</p>
          <p><strong>URL:</strong> <a href={paper.url} className="text-blue-600 underline" target="_blank">{paper.url}</a></p>
          <p><strong>PDF:</strong> <a href={paper.pdf_url} className="text-blue-600 underline" target="_blank">{paper.pdf_url || '未入力'}</a></p>
          <p><strong>引用数:</strong> {paper.citation_count === -1 ? '未入力' : paper.citation_count}</p>
          <p><strong>BibTeX:</strong></p>
          <pre className="bg-gray-100 p-2 rounded text-xs overflow-x-auto">{paper.bibtex || '未入力'}</pre>
        </div>
      </div>
    </div>
  );
}
