import { title } from "@/components/primitives";
import DefaultLayout from "@/layouts/default";
import {Chip} from "@heroui/chip";

export default function DocsPage() {
  return (
    <DefaultLayout>
      <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
        <div className="inline-block max-w-lg text-center justify-center">
          <span className={title({size:"sm", color:"cyan"})}>About&nbsp;</span>
        </div>
        <div className="row marketing">
            <div className="inline-block max-w-lg text-center justify-center">
                <Chip>What is OMCI Analyzer?</Chip>
                <p>OMCI Analyzer is a OMCI model auto generate tool</p>
                <Chip>What is file type supported?</Chip>
                <p>
                ISAM OMCI log, LightSpan engine log, Nokia ONU log, GponDoctor CSV log
                </p>
                <Chip>Submit bugs or can't parse log?</Chip>
                <p>
                please submit bugs and log in conflunce{" "}
                <a
                    href="https://confluence.ext.net.nokia.com/display/Transpt2/Troubleshooting+-+OMCI+Analyzer"
                    target="_blank"
                    className="text-blue-600 underline"
                >
                    OMCI Analyzer
                </a>
                </p>
                <Chip>Support</Chip>
                <p>
                Please contact Transport2 eonuMgnt team or send mail to minghe.huang@nokia.com
                </p>
                <Chip>What's the parse speed?</Chip>
                <p>
                base on the log file size, will take couple of seconds to tens seconds
                </p>
                <Chip>Why parse failed?</Chip>
                <p>Unkonw ME or unadapted log formart</p>
                <Chip>Should keep web page after parse done?</Chip>
                <p>Needn't, you also can check your log result in present models page</p>
            </div>
        </div>
      </section>
    </DefaultLayout>
  );
}
