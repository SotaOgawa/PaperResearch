// frontend/__tests__/index.test.tsx
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import HomePage from "../pages/index";
import * as api from "../lib/api";
import '@testing-library/jest-dom';

// モックデータ
const mockPapers = [
  {
    id: 1,
    title: "Transformer is All You Need",
    authors: "Taro Yamada, Hanako Suzuki",
    conference: "NeurIPS",
    year: 2017,
    url: "https://arxiv.org/abs/1706.03762",
  },
];

describe("HomePage 論文検索機能", () => {
  it("検索ボタンでfetchJSONが呼ばれ、論文結果が表示される", async () => {
    // fetchJSONをモック
    jest.spyOn(api, "fetchJSON").mockResolvedValueOnce({ papers: mockPapers });

    render(<HomePage />);

    // 入力
    fireEvent.change(screen.getByPlaceholderText("タイトルで検索"), {
      target: { value: "Transformer" },
    });

    fireEvent.change(screen.getByPlaceholderText("発表年"), {
      target: { value: "2017" },
    });

    // 検索ボタンを押す
    fireEvent.click(screen.getByText("検索"));

    // 結果が出るのを待つ
    await waitFor(() =>
      expect(screen.getByText("Transformer is All You Need")).toBeInTheDocument()
    );

    expect(screen.getByText("Taro Yamada, Hanako Suzuki")).toBeInTheDocument();
    expect(screen.getByText("NeurIPS / 2017")).toBeInTheDocument();
    expect(screen.getByText("論文リンク")).toHaveAttribute("href", expect.stringContaining("arxiv"));
  });
});