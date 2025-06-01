import React, {ReactElement} from "react";
import { render, screen } from "@testing-library/react";
import '@testing-library/jest-dom';

function DummyComponent() {
  return <div>テスト表示</div>;
}

describe("コンポーネントの表示", () => {
  it("コンポーネントが正しく表示される", () => {
    render(<DummyComponent />);
    expect(screen.getByText("テスト表示")).toBeInTheDocument();
  });
});