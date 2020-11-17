import React from "react";
import Prism from "prismjs";

export const CodeSnippet = ({ language, children }) => (
  <pre className={`language-${language}`}>
    <code
      className={`language-${language}`}
      dangerouslySetInnerHTML={{
        __html: Prism.highlight(children, Prism.languages[language], language),
      }}
    />
  </pre>
);
